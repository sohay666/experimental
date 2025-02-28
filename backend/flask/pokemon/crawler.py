import requests
import argparse
import csv
import json
import re
from bs4 import BeautifulSoup

URL = 'https://pokemondb.net/pokedex/national'

def save(output_file, pokemon):
    fieldnames=['id', 'name', 'generation', 'types', 'species','height', 'weight', 'abilities', 'image']
    with open(output_file, mode='a', newline='', encoding='utf-8') as file:
        writer = csv.DictWriter(file, fieldnames=fieldnames)

        # Write the header if the file is empty
        file.seek(0, 2)  # Move cursor to the end of the file
        if file.tell() == 0:
            writer.writeheader()  # Write header if the file is empty

        # Write Pokémon data to the CSV
        writer.writerow({
            'id': pokemon['id'],
            'name': pokemon['name'],
            'generation': pokemon['generation'],
            'types': pokemon['types'],
            'species': pokemon['species'],
            'height': pokemon['height'],
            'weight': pokemon['weight'],
            'abilities': pokemon['abilities'],
            'image': pokemon['image'],
        })


def getDetailPokemon(URL):
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
        "Cookie": "hb_insticator_uid=3d9af732-208d-49a5-a128-99722a15b1f1; _cc_id=5e1d1d518dae8c28ad3c6c5f673b026e; panoramaId_expiry=1738399141541; panoramaId=805cd1bfefdc9b6df0301e5893254945a702f054c706bb25f86191a43232d718; panoramaIdType=panoIndiv; __gads=ID=33e60a08e7879321:T=1737794342:RT=1737794342:S=ALNI_MapVHYsUNESGOQiDtoXz2XSSnHhZQ; __gpi=UID=0000100d1fbe8b2e:T=1737794342:RT=1737794342:S=ALNI_MYxNSCQ2xh5SjT-8jJrjYRp_OPXPA; __eoi=ID=ca56db7539d2fd21:T=1737794342:RT=1737794342:S=AA-AfjYMlEZGCpzym-a28bPhpFDh",
        "Accept": "*/*",
        "Connection": "keep-alive"
    }
    response = requests.get(URL, headers=headers, timeout=15, verify=True)
    html_data = response.text
    # Parse the HTML
    soup = BeautifulSoup(html_data, 'html.parser')

    # Extract Type
    types = [type_tag.text.strip() for type_tag in soup.select('tr:has(th:-soup-contains("Type")) td a')]

    # Extract Species
    species = soup.select_one('tr:has(th:contains("Species")) td').text.strip()

    # Extract Height
    height = soup.select_one('tr:has(th:contains("Height")) td').text.strip()

    # Extract Weight
    weight = soup.select_one('tr:has(th:contains("Weight")) td').text.strip()

    # Extract Abilities
    abilities = [ability.text.strip() for ability in soup.select('tr:has(th:-soup-contains("Abilities")) td a')]
    
    gen_type = "1"
    gen_types = html_data.split("Generation ")
    if len(gen_types) > 0:
        gen_type = gen_types[1].split("</abbr>")[0]


    output = {  
        'generation': gen_type,
        'types': json.dumps(types), 
        'species': species,
        'height': height,
        'weight': weight,
        'abilities': json.dumps(abilities),
    }
    return output


def getListPokemon(args):
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
        "Cookie": "hb_insticator_uid=3d9af732-208d-49a5-a128-99722a15b1f1; _cc_id=5e1d1d518dae8c28ad3c6c5f673b026e; panoramaId_expiry=1738399141541; panoramaId=805cd1bfefdc9b6df0301e5893254945a702f054c706bb25f86191a43232d718; panoramaIdType=panoIndiv; __gads=ID=33e60a08e7879321:T=1737794342:RT=1737794342:S=ALNI_MapVHYsUNESGOQiDtoXz2XSSnHhZQ; __gpi=UID=0000100d1fbe8b2e:T=1737794342:RT=1737794342:S=ALNI_MYxNSCQ2xh5SjT-8jJrjYRp_OPXPA; __eoi=ID=ca56db7539d2fd21:T=1737794342:RT=1737794342:S=AA-AfjYMlEZGCpzym-a28bPhpFDh",
        "Accept": "*/*",
        "Connection": "keep-alive"
    }

    response = requests.get(URL, headers=headers,timeout=10, verify=True)

    html_data = response.text

    # Parse the HTML
    soup = BeautifulSoup(html_data, 'html.parser')

    # Find all infocards
    pokemon_list = []
    infocards = soup.find_all('div', class_='infocard')
    for card in infocards:
        pokemon_id = card.find('small').text.strip()
        pokemon_name = card.find('a', class_='ent-name').text.strip()
        pokemon_link = 'https://pokemondb.net'+card.find('a')['href']
        pokemon_image = card.find('img', class_='img-fixed')['src']
        pokemon_type = card.find('a', class_='itype').text.strip()
        
        # Store the data in a dictionary
        pokemon_list.append({
            'id': pokemon_id,
            'link': pokemon_link,
            'name': pokemon_name,
            'image': pokemon_image,
            'type': pokemon_type,
        })

    # Print the extracted data
    for pokemon in pokemon_list:
        details = getDetailPokemon(pokemon['link'])
        pokemon['generation'] = details['generation']
        pokemon['types'] = details['types']
        pokemon['species'] = details['species']
        pokemon['height'] = details['height']
        pokemon['weight'] = details['weight']
        pokemon['abilities'] = details['abilities']
        print(f"generation: {pokemon['generation']}, ID: {pokemon['id']}, Name: {pokemon['name']}, height: {pokemon['height']}, weight: {pokemon['weight']}, abilities: {pokemon['abilities']}, Species: {pokemon['species']},Types: {pokemon['types']}, Image URL: {pokemon['image']}, Link: {pokemon['link']}")
        save(args.output, pokemon)

def main(args):
    getListPokemon(args)
    
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Crawler data pokemon.")
    parser.add_argument("-o", "-o", dest='output', default="pokemon_data.csv", help="Output csv file name, e.g: pokemon_data.csv")
    args = parser.parse_args()
    main(args)

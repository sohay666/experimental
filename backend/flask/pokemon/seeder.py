import csv
import json
import os
import argparse
import mysql.connector
from mysql.connector import Error
from dotenv import load_dotenv
load_dotenv()  # This will load environment variables from a .env file, if it exists

# Function to connect to the database
def connect_db():
    try:
        connection = mysql.connector.connect(
            host=os.getenv('MYSQL_HOST'),  # replace with your MySQL host
            database=os.getenv('MYSQL_DATABASE'),  # replace with your database name
            user=os.getenv('MYSQL_USER'),  # replace with your MySQL user
            password=os.getenv('MYSQL_PASSWORD')  # replace with your MySQL password
        )
        if connection.is_connected():
            print("Connected to MySQL database")
            return connection
    except Error as e:
        print(f"Error while connecting to MySQL: {e}")
        return None

# Function to insert Pokemon data into MySQL
def insert_monster(data):
    connection = connect_db()
    if connection is None:
        print("Failed to connect to database")
        return

    try:
        cursor = connection.cursor()
        query = """
            INSERT INTO monsters (name, generation, types, species, height, weight, abilities, image, created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
        """
        # Debugging: print the actual values you're passing to the query
        print(f"Executing query with values: {data['name']}, {data['generation']}, {json.dumps(data['types'])}, "
              f"{data['species']}, {data['height']}, {data['weight']}, {json.dumps(data['abilities'])}, {data['image']}")

        cursor.execute(query, (data['name'], data['generation'], json.dumps(data['types']), data['species'], data['height'], data['weight'],
                               json.dumps(data['abilities']),  data['image']))
        connection.commit()
        print(f"Inserted {data['name']} into database")
    except Error as e:
        print(f"Error while inserting data: {e}")
    finally:
        cursor.close()
        connection.close()

# Function to load Pokémon data from CSV
def load_csv_data(csv_file):
    pokemon_list = []
    
    # Open the CSV file and read its content
    with open(csv_file, mode='r', newline='', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        for row in reader:
            # Convert JSON strings back to Python objects for abilities and evolution
            row['abilities'] = json.loads(row['abilities'])
            row['types'] = json.loads(row['types'])
            pokemon_list.append(row)
    
    return pokemon_list

# Main function to load CSV data and insert into database
def load_and_insert(csv_file):
    # Load data from CSV
    pokemon_list = load_csv_data(csv_file)
    
    # Insert each Pokemon's data into the MySQL database
    for pokemon in pokemon_list:
        insert_monster(pokemon)

def main(args):
    load_and_insert(args.file)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Data seeder for pokemon into your database.")
    parser.add_argument("-f", "-f", dest='file', default="pokemon_data.csv", help="File csv file name, e.g: pokemon_data.csv")
    args = parser.parse_args()
    main(args)

import os
import json
from flask import Flask, request, jsonify

import mysql.connector
# Optionally load the .env file for local development
from dotenv import load_dotenv
load_dotenv()  # This will load environment variables from a .env file, if it exists

app = Flask(__name__)

def connect_db():
    try:
        connection = mysql.connector.connect(
            host=os.getenv('MYSQL_HOST'),
            database=os.getenv('MYSQL_DATABASE'),
            user=os.getenv('MYSQL_USER'),
            password=os.getenv('MYSQL_PASSWORD'),
        )
        if connection.is_connected():
            return connection
    except Exception as e:
        print(f"Error: {e}")
        return None

# Function to get Pokémon data with pagination and filtering
def get_pokemon_data(page=1, filters=None):
    connection = connect_db()
    if connection is None:
        return []
    
    cursor = connection.cursor(dictionary=True)
    filter_by_name = filters["name"] if filters != None else None
    filter_by_generation = filters["generation"] if filters != None else None

    # Setting up the filter part
    params = []
    filter_sql = ''
    # Adding filter for name if present
    if filter_by_name:
        filter_sql += "WHERE name LIKE %s"
        filter_by_name = f"%{filter_by_name}%"
        params.append(filter_by_name)

    # Adding filter for generation if present
    if filter_by_generation:
        if filter_sql:
            filter_sql += " AND generation = %s"
        else:
            filter_sql = "WHERE generation = %s"
        params.append(filter_by_generation)

    limit = 10
    offset = (page - 1) * limit
    
    query = f"""
        SELECT * FROM monsters
        {filter_sql}
        ORDER BY name
        LIMIT %s OFFSET %s
    """
    
    # Adding the pagination limit and offset
    params.extend([limit, offset])

    cursor.execute(query, tuple(params))

    # cursor.execute(query, (filter_query, filter_query, limit, offset) if filter_query else (limit, offset))
    pokemon_data = cursor.fetchall()
    
    # Convert JSON strings to Python objects
    for pokemon in pokemon_data:
        if 'abilities' in pokemon and pokemon['abilities']:
            pokemon['abilities'] = json.loads(pokemon['abilities'])
        if 'types' in pokemon and pokemon['types']:
            pokemon['types'] = json.loads(pokemon['types'])

    cursor.close()
    connection.close()
    
    return pokemon_data

# Function to get the total count of Pokémon for pagination
def get_total_pokemon_count(filters=None):
    connection = connect_db()
    if connection is None:
        return 0
    
    cursor = connection.cursor()
    
    filter_by_name = filters["name"] if filters != None else None
    filter_by_generation = filters["generation"] if filters != None else None

    # Setting up the filter part
    params = []
    filter_sql = ''
    # Adding filter for name if present
    if filter_by_name:
        filter_sql += "WHERE name LIKE %s"
        filter_by_name = f"%{filter_by_name}%"
        params.append(filter_by_name)

    if filter_by_generation:
        if filter_sql:
            filter_sql += " AND generation = %s"
        else:
            filter_sql = "WHERE generation = %s"
        params.append(filter_by_generation)
    
    query = f"SELECT COUNT(*) FROM monsters {filter_sql}"
    cursor.execute(query, tuple(params))

    total_count = cursor.fetchone()[0]
    
    cursor.close()
    connection.close()
    
    return total_count

@app.route("/")
def index():
    return jsonify({"message": "Success", "status_code": 200})

@app.route("/api/v1/pokemon")
def pokemon():
    page = request.args.get('page', 1, type=int)
    filters = {
        'name': request.args.get('name', '', type=str),
        'generation': request.args.get('gen', '', type=int)
    }
    
    # Fetch Pokémon data with pagination and filter
    pokemon_data = get_pokemon_data(page, filters)
    total_count = get_total_pokemon_count(filters)
    
    # Calculate total pages
    total_pages = (total_count // 10) + (1 if total_count % 10 else 0)
    
    # Prepare the response
    response = {
        "message": "Success",
        "status_code": 200,
        "data": {
            "pokemon": pokemon_data,
            "page": page,
            "total_pages": total_pages,
            "total_count": total_count,
        }
    }
    return jsonify(response)

if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=8888)

// mongo-init/init.js

// Switch to the DB you want to use
db = db.getSiblingDB("magic-stream");

// Create collections
db.createCollection("genres");
db.createCollection("movies");
db.createCollection("rankings");
db.createCollection("users");
db.createCollection("user_sessions")

// Optional: insert a sample document to test
db.users.insertMany([
    {
        "user_id": "68385b9981097c6b4042dab4",
        "first_name": "Bob",
        "last_name": "Jones",
        "email": "bobjones@hotmail.com",
        "role": "ADMIN",
        "created_at": new Date("2025-05-29T13:05:29.000Z"),
        "updated_at": new Date("2025-05-29T13:06:51.000Z"),
        "favourite_genres": [
            {"genre_id": 1, "genre_name": "Comedy"},
            {"genre_id": 4, "genre_name": "Fantasy"}
        ]
    },
    {
        "user_id": "684535a8c3d7e6b4ac1c5203",
        "first_name": "Sarah",
        "last_name": "Smith",
        "email": "sarahsmith@hotmail.com",
        "role": "USER",
        "created_at": new Date("2025-06-08T07:03:04.865Z"),
        "updated_at": new Date("2025-06-23T09:26:11.000Z"),
        "favourite_genres": [
            {"genre_id": 5, "genre_name": "Thriller"},
            {"genre_id": 6, "genre_name": "Sci-fi"},
            {"genre_id": 8, "genre_name": "Mystery"}
        ]
    },
    {
        "user_id": "68512bc5a29d3bfe81637e3f",
        "first_name": "Ben",
        "last_name": "Madison",
        "email": "benmadison@hotmail.com",
        "role": "USER",
        "created_at": new Date("2025-06-17T08:48:05.649Z"),
        "updated_at": new Date("2025-06-17T08:48:05.649Z"),
        "favourite_genres": [
            {"genre_id": 1, "genre_name": "Comedy"},
            {"genre_id": 5, "genre_name": "Thriller"},
            {"genre_id": 6, "genre_name": "Sci-Fi"}
        ]
    }
]);

db.rankings.insertMany([
    {
        "ranking_value": 999,
        "ranking_name": "Not_Ranked"
    },
    {
        "ranking_value": 1,
        "ranking_name": "Excellent"
    },
    {
        "ranking_value": 2,
        "ranking_name": "Good"
    },
    {
        "ranking_value": 3,
        "ranking_name": "Okay"
    },
    {
        "ranking_value": 4,
        "ranking_name": "Bad"
    },
    {
        "ranking_value": 5,
        "ranking_name": "Terrible"
    }
]);

db.movies.insertMany([
    {
        "imdb_id": "tt0111161",
        "title": "The Shawshank Redemption",
        "poster_path": "https://image.tmdb.org/t/p/w300/2GgerXCbCMgvt2kLwWEmJWCSG65.jpg",
        "youtube_id": "PLl99DlL6b4",
        "genre": [
            {
                "genre_id": 2,
                "genre_name": "Drama"
            }
        ],
        "admin_review": "I loved the acting in this movie. It was absolutely sublime!",
        "ranking": {
            "ranking_value": 1,
            "ranking_name": "Excellent"
        }
    },
    {
        "imdb_id": "tt7131622",
        "title": "Once upon a time in Hollywood",
        "poster_path": "https://image.tmdb.org/t/p/w300/wQKeS2JrsRF8XSfd9zqflrc5gad.jpg",
        "youtube_id": "ELeMaP8EPAA",
        "genre": [
            {
                "genre_id": 2,
                "genre_name": "Drama"
            },
            {
                "genre_id": 1,
                "genre_name": "Comedy"
            }
        ],
        "admin_review": "This movie is awful.",
        "ranking": {
            "ranking_name": "Terrible",
            "ranking_value": 5
        }
    },
    {
        "imdb_id": "tt0080339",
        "title": "Airplane!",
        "poster_path": "https://image.tmdb.org/t/p/w300/zOiB3p2WTTiwCFgTMnXuDGgzbTN.jpg",
        "youtube_id": "07pPmCfKi3U",
        "genre": [
            {
                "genre_id": 1,
                "genre_name": "Comedy"
            }
        ],
        "admin_review": "I didn't love this movie but I didn't hate it either.",
        "ranking": {
            "ranking_value": 3,
            "ranking_name": "Okay"
        }
    },
    {
        "imdb_id": "tt1119646",
        "title": "The Hangover",
        "poster_path": "https://image.tmdb.org/t/p/w300/c15rH9S5JN83UXqu9iM4TQsW6Rl.jpg",
        "youtube_id": "jj6wcUes1no",
        "genre": [
            {
                "genre_id": 1,
                "genre_name": "Comedy"
            }
        ],
        "admin_review": "I didn't love this movie but I didn't hate it either.",
        "ranking": {
            "ranking_value": 3,
            "ranking_name": "Okay"
        }
    },
    {
        "imdb_id": "tt0109040",
        "title": "Ace Ventura: Pet Detective",
        "poster_path": "https://image.tmdb.org/t/p/w300/dukxJWd72ffNWfqfFSVFFuym4RG.jpg",
        "youtube_id": "qjBb1CKLpzE",
        "genre": [
            {
                "genre_id": 1,
                "genre_name": "Comedy"
            }
        ],
        "admin_review": "This is really not so great.",
        "ranking": {
            "ranking_value": 4,
            "ranking_name": "Bad"
        }
    },
    {
        "imdb_id": "tt0105695",
        "title": "Unforgiven",
        "poster_path": "https://image.tmdb.org/t/p/w300/7W0CsZBe5fkX29DvHBi5Ct1WqLe.jpg",
        "youtube_id": "6_UlfsdGiEc",
        "genre": [
            {
                "genre_id": 2,
                "genre_name": "Drama"
            },
            {
                "genre_id": 3,
                "genre_name": "Western"
            }
        ],
        "admin_review": "I hate this movie",
        "ranking": {
            "ranking_value": 5,
            "ranking_name": "Terrible\n"
        }
    },
    {
        "imdb_id": "tt0060196",
        "title": "The Good, The Bad, and the Ugly",
        "poster_path": "https://image.tmdb.org/t/p/w300/e881nA7p982CHL5GjI1LICwHMd7.jpg",
        "youtube_id": "WCN5JJY_wiA",
        "genre": [
            {
                "genre_id": 3,
                "genre_name": "Western"
            }
        ],
        "admin_review": "I did not like this!",
        "ranking": {
            "ranking_value": 4,
            "ranking_name": "Bad"
        }
    },
    {
        "imdb_id": "tt0903624",
        "title": "The Hobbit: An Unexpected Journey",
        "poster_path": "https://image.tmdb.org/t/p/w300/vdAGcr1F6wJPRryeODVAcy2mU4z.jpg",
        "youtube_id": "9PSXjr1gbjc",
        "genre": [
            {
                "genre_id": 4,
                "genre_name": "Fantasy"
            }
        ],
        "admin_review": "The movie was aweful! I really hated it.",
        "ranking": {
            "ranking_value": 5,
            "ranking_name": "Terrible"
        }
    },
    {
        "imdb_id": "tt0241527",
        "title": "Harry Potter and the Philosopher's Stone",
        "poster_path": "https://image.tmdb.org/t/p/w300/e6JYlushXIXK85JGfDHEFHrrNYK.jpg",
        "youtube_id": "VyHV0BRtdxo",
        "genre": [
            {
                "genre_id": 4,
                "genre_name": "Fantasy"
            }
        ],
        "admin_review": "This movie wasn't great but it wasn't bad either.",
        "ranking": {
            "ranking_value": 3,
            "ranking_name": "Okay"
        }
    },
    {
        "imdb_id": "tt2267998",
        "title": "Gone Girl",
        "poster_path": "https://image.tmdb.org/t/p/w300/xpA0q0DJWKe7AY63pVPZbGLwuo5.jpg",
        "youtube_id": "2-_-1nJf8Vg",
        "genre": [
            {
                "genre_id": 2,
                "genre_name": "Drama"
            },
            {
                "genre_id": 4,
                "genre_name": "Fantasy"
            }
        ],
        "admin_review": "An okay but dark movie.",
        "ranking": {
            "ranking_value": 3,
            "ranking_name": "Okay"
        }
    },
    {
        "imdb_id": "tt8946378",
        "title": "Knives Out",
        "poster_path": "https://image.tmdb.org/t/p/w300/hVcCOlHU0HmGBBQNmS9RlalBXGz.jpg",
        "youtube_id": "qGqiHJTsRkQ",
        "genre": [
            {
                "genre_id": 2,
                "genre_name": "Drama"
            },
            {
                "genre_id": 5,
                "genre_name": "Thriller"
            },
            {
                "genre_id": 8,
                "genre_name": "Mystery"
            }
        ],
        "admin_review": "The story was fantastic and the acting was sublime!",
        "ranking": {
            "ranking_value": 1,
            "ranking_name": "Excellent"
        }
    },
    {
        "imdb_id": "tt0080684",
        "title": "Star Wars: The Empire Strikes Back",
        "poster_path": "https://image.tmdb.org/t/p/w300/1mh82R1qLKwdutQVGnpHItdwPCP.jpg",
        "youtube_id": "JNwNXF9Y6kY",
        "genre": [
            {
                "genre_id": 6,
                "genre_name": "Sci-Fi"
            },
            {
                "genre_id": 4,
                "genre_name": "Fantasy"
            }
        ],
        "admin_review": "This is okay!",
        "ranking": {
            "ranking_name": "Okay",
            "ranking_value": 3
        }
    },
    {
        "imdb_id": "tt0102975",
        "title": "Star Trek: The Undiscovered Country",
        "poster_path": "https://image.tmdb.org/t/p/w300/jukY1tFpuXgqrJLl1PvdOMarCvN.jpg",
        "youtube_id": "RYA2q2Sm_Jo",
        "genre": [
            {
                "genre_id": 6,
                "genre_name": "Sci-Fi"
            },
            {
                "genre_id": 7,
                "genre_name": "Action"
            }
        ],
        "admin_review": "It wasn't good but it wasn't too bad either.",
        "ranking": {
            "ranking_value": 3,
            "ranking_name": "Okay"
        }
    },
    {
        "imdb_id": "tt1297919",
        "title": "Blitz",
        "poster_path": "https://image.tmdb.org/t/p/w300/tI4fPu0LZ0FNdZdi6fvYGCRiuQs.jpg",
        "youtube_id": "mhO2WJ3MNRI",
        "genre": [
            {
                "genre_id": 9,
                "genre_name": "Crime"
            },
            {
                "genre_id": 7,
                "genre_name": "Action"
            },
            {
                "genre_id": 5,
                "genre_name": "Thriller"
            }
        ],
        "admin_review": "This was a lovely movie!",
        "ranking": {
            "ranking_value": 1,
            "ranking_name": "Excellent"
        }
    },
    {
        "imdb_id": "tt0790724",
        "title": "Jack Reacher",
        "poster_path": "https://image.tmdb.org/t/p/w300/8sih1uieUopA9zrEO1meNhmm7aO.jpg",
        "youtube_id": "A7FiWkyevqY",
        "genre": [
            {
                "genre_id": 7,
                "genre_name": "Action"
            },
            {
                "genre_id": 5,
                "genre_name": "Thriller"
            }
        ],
        "admin_review": "This might be the worst movie that I have ever seen in my life!",
        "ranking": {
            "ranking_value": 5,
            "ranking_name": "Terrible"
        }
    }
]);

db.genres.insertMany([
        {
            "genre_id": 1,
            "genre_name": "Comedy"
        },
        {
            "genre_id": 2,
            "genre_name": "Drama"
        },
        {
            "genre_id": 3,
            "genre_name": "Western"
        },
        {
            "genre_id": 4,
            "genre_name": "Fantasy"
        },
        {
            "genre_id": 5,
            "genre_name": "Thriller"
        },
        {
            "genre_id": 6,
            "genre_name": "Sci-Fi"
        },
        {
            "genre_id": 7,
            "genre_name": "Action"
        },
        {
            "genre_id": 8,
            "genre_name": "Mystery"
        },
        {
            "genre_id": 9,
            "genre_name": "Crime"
        }
    ]);
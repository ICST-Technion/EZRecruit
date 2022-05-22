from tokenize import String
from flask import Flask
from flask_restful import Resource, Api, reqparse
from pdf_parser_core import searchCVsInFolder
import json


app = Flask(__name__)
api = Api(app)


class CVs(Resource):
    """def get(self):
        parser = reqparse.RequestParser()
        parser.add_argument('pathToDir',required=False)
        parser.add_argument('wordsList',required=False)
        args = parser.parse_args()
        print(args)
        return {},200"""

    """
        get a list of cvs...
    """
    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument('pathToDir',required=True)
        parser.add_argument('wordsList',required=True)

        args = parser.parse_args()
        # print(args)
        # print(args['pathToDir'])
        # print(args['wordsList'])
        path = args['pathToDir']
        # print(path)
        wordsList = []

        wordsListStr = args['wordsList']
        wordsListPairs = wordsListStr[1:-1].split(',')
        # print(wordsListPairs)

        for pair in wordsListPairs:
            word:String = pair.split(':')[0].strip()[1:-1]
            wordsList.append(word)

        cvsList = searchCVsInFolder(path, wordsList)
        
        return json.dumps(cvsList, ensure_ascii=False), 200
        


api.add_resource(CVs, '/cvs')  # '/users' is our entry point


if __name__ == '__main__':
    app.run()  # run our Flask app
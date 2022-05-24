from multilingual_pdf2text.pdf2text import PDF2Text
from multilingual_pdf2text.models.document_model.document import Document
import re
import os

lng = "heb" # the language of the CV's, now the default is hebrew

"""
    core functions
    Single CV:
    - check if a specific (single) word is in CV
    - check if multiple words are in CV - return number of words found
    - TODO: return mapping of words in a CV {word: numOfAppearances}

    Directory:
    - check all CV's in the directory for words
"""

"""
    Concat the strings from a given pdf file:
    @param content: list of pages data

    @return: a single string contains all pages data
"""
def concatPagesStrings(content):
    wholeFileString = ""
    for page in content:
        wholeFileString += page['text']
    return wholeFileString


"""
    Extract data from a pdf file
    @param filePath: the path for the file to

    @return: the extracted data as string
"""
def extractDataFromFile(filePath):
    if os.path.isfile(filePath) and filePath.endswith(".pdf"):
        pdf_document = Document(
            document_path=filePath,
            language=lng
            )
        pdf2text = PDF2Text(document=pdf_document)
        content = pdf2text.extract() # get the content of all pages in the file: list of dictionaries [{}]
        data = concatPagesStrings(content)
        return data
    else: 
        return None


"""
    Tries to find a word in a CV:
    @param word: the word to search

    @return: regex function
"""
def findSpecificWord(word):
    return re.compile(r'\b({0})\b'.format(word), flags=re.IGNORECASE).search


"""
    Find a batch of words in a given CV
    @param wordsList [List[string]]: List of words to search for
    @param cv [string]: CV to search 

    @return hitsCount [int]: Number of hits
            hitWords [dict(word -> Boolean)]: words and a boolean indicating if there was a match
"""
def findWords(wordsList, cv):
    hitsCount = 0
    hitWords = {}
    for word in wordsList:
        hitWords[word] = False
        result = findSpecificWord(word)(cv)
        if (result != None):
            hitsCount += 1
            hitWords[word] = True
    return hitsCount, hitWords


"""
    Search for matching words in all files (CV's) in a given directory
    @param folderPath: the path (TODO: check if global only) of the directory
    @param wordsList: a wordsList of words to search
    
    @return: a list of tuples (file, hitsCount, hitWords) sorted by hitsCount backwords
"""
def searchCVsInFolder(folderPath, wordsList):
    resultsList = []

    # iterate over the Cv's in folder
    for filename in os.listdir(folderPath):
        f = os.path.join(folderPath, filename)
        # checking if it's a pdf file (a CV)
        data = extractDataFromFile(f)
        if data is None:
            continue
        hitsCount, hitWords = findWords(wordsList, data)
        resultsList.append((filename, hitsCount, hitWords))
    
    # sort the list by hitsCount
    resultsList.sort(key=lambda tup: tup[1], reverse=True)
    return resultsList


"""
    Search for matching words in a given file (CV)
    @param filePath: the path of the file
    @param wordsList: a wordsList of words to search
    
    @return: num of hits and a dict of words:boolean indicating match results 
"""
def searchSingleCv(filePath, wordsList):
    data = extractDataFromFile(filePath)
    if data is None:
        return None
    fileHead, fileName = os.path.split(filePath)
    hitsCount, hitWords = findWords(wordsList, data)
    return [[fileName, hitsCount, hitWords]]



"""
TODO: this method is stil very basic and should be improved (using regex and more...)

    Return mapping of words to number of appearances in the file
    @param filePath: the path of the file
    
    @return: a dictionary mapping words to number of appearances 
"""
def getMappingOfWordsInCV(filePath):
    data = extractDataFromFile(filePath)
    if data is None:
        return None
    wordMapping = {}
    data = data.replace('\n',' ').split(' ')
    for word in data:
        if word in wordMapping:
            counter = int(wordMapping[word])
            wordMapping[word] = counter + 1
        else:
            wordMapping[word] = 1
    
    # print(wordMapping)
    return wordMapping



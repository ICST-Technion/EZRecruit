import errno
import requests
import json
from fileinput import FileInput
import PySimpleGUI as sg
import os.path

# sg.theme_previewer()
sg.theme('BlueMono')    # Keep things interesting for your users

API_ENDPOINT = "http://localhost:5000/cvs"


def buildJsonForSingleFileReq(filePath, wordsListString):
    jsonForReq = {}
    jsonForReq["pathToDir"] = filePath
    wordsList = wordsListString.split(',')
    wordsJson = {}
    for word in wordsList:
        wordsJson[word] = "1"
    jsonForReq["wordsList"] = wordsJson
    return jsonForReq

def printResults(outputWindow, data):
    fullString = ""
    for cv in data:
        fullString += "Name: " + cv[0] + "\n"
        fullString += "Wordcount: " + str(cv[1]) + "\n"
        fullString += "Description: " + str(cv[2]) + "\n"
        fullString += "----------------------------------------------------------------\n"
    outputWindow.update(fullString)



def dirSearchWindow():
    file_list_column = [
    [
        sg.Text("CV's Folder"),
        sg.In(size=(25, 1), enable_events=True, key="-FOLDER-"),
        sg.FolderBrowse(),
    ],
    [
        sg.Listbox(
            values=[], enable_events=True, size=(40, 20), key="-FILE LIST-"
        )
    ],
]

    image_viewer_column = [
        [sg.Text("Choose an image from list on left:")],
        [sg.Text(size=(40, 1), key="-TOUT-")],
        [sg.Image(key="-IMAGE-")],
    ]


    layout = [
        [
            sg.Column(file_list_column),
            sg.VSeperator(),
            sg.Column(image_viewer_column)
        ]
    ]
    window = sg.Window(title="Pdf Parser", layout=layout)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break

        if event == "-FOLDER-":
            folder = values["-FOLDER-"]
            try:
                # Get list of files in folder
                file_list = os.listdir(folder)
            except:
                file_list = []

            fnames = [
                f
                for f in file_list
                if os.path.isfile(os.path.join(folder, f))
                and f.lower().endswith((".pdf"))
            ]
            window["-FILE LIST-"].update(fnames)
        elif event == "-FILE LIST-":  # A file was chosen from the listbox
            pass
            """try:
                filename = os.path.join(
                    values["-FOLDER-"], values["-FILE LIST-"][0]
                )
                window["-TOUT-"].update(filename)
                window["-IMAGE-"].update(filename=filename)
            except:
                print("whoops")
                pass"""

    window.close()

def fileSearchWindow():
    file_list_column = [
    [
        sg.Text("CV's Folder"),
        sg.In(size=(25, 1), enable_events=True, key="-FOLDER-"),
        sg.FolderBrowse(),
    ],
    [
        sg.Text('Words List'),
        # sg.In(size=(25, 1), enable_events=True, key="-WordsList-"),
        sg.InputText(size=(25, 1), enable_events=True, key="-WordsList-"),
        # sg.Button("Button", key="-PRESSED-")
    ],
    [
        sg.Listbox(
            values=[], enable_events=True, size=(40, 20), key="-FILE LIST-"
        )
    ],
    [
        sg.Column([[sg.Button("Scan", key="-SCAN-")]], justification='center')
    ]
]

    image_viewer_column = [
        [sg.Text("Choose a file from list on left:")],
        [sg.Text(size=(40, 1), key="-TOUT-")],
        [sg.Text(key="CV-RESULTS")],
        [sg.Multiline(size=(50,10), key="-OUTPUT-")]
    ]


    layout = [
        [
            sg.Column(file_list_column),
            sg.VSeperator(),
            sg.Column(image_viewer_column)
        ]
    ]
    window = sg.Window(title="Pdf Parser", layout=layout)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break

        if event == "-FOLDER-":
            folder = values["-FOLDER-"]
            try:
                # Get list of files in folder
                file_list = os.listdir(folder)
            except:
                file_list = []

            fnames = [
                f
                for f in file_list
                if os.path.isfile(os.path.join(folder, f))
                and f.lower().endswith((".pdf"))
            ]
            window["-FILE LIST-"].update(fnames)
        elif event == "-FILE LIST-":  # A file was chosen from the listbox
            pass
            try:
                filename = os.path.join(
                    values["-FOLDER-"], values["-FILE LIST-"][0]
                )
                window["-TOUT-"].update(filename)
                # window["-IMAGE-"].update(filename=filename)
            except:
                print("whoops")
                pass
        elif event == "-SCAN-":
            #print(values["-WordsList-"])
            try:
                filePath = filename
                # print("filePath: " + filePath)
                wordsList = values["-WordsList-"]
                # print("wordsList: " + wordsList)
                jsonForReq = buildJsonForSingleFileReq(filePath, wordsList)
                r = requests.post(url = API_ENDPOINT, json = jsonForReq, headers={'content-type': 'application/json'})
                data = json.loads(r.json())
                # print(data)
                printResults(window["-OUTPUT-"], data)
                # window["-OUTPUT-"].print(data)
            except:
                print("ERROR IN PARSING")
            

    window.close()


def mainWindow():
    layout = [
        [
            sg.Button("Search Single File", key="openFileSearch"),
            sg.Button("Search Directory", key="openDirectorySearch")

        ]
    ]
    window = sg.Window(title="Main Window", layout=layout)
    while True:
        event, values = window.read()
        # End program if user closes window or
        # presses the OK button
        if event == "EXIT" or event == sg.WIN_CLOSED:
            break
        elif event == "openDirectorySearch":
            dirSearchWindow()
        elif event == "openFileSearch":
            fileSearchWindow()

if __name__ == "__main__":
    mainWindow()
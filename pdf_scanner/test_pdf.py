import os
import codecs
import glob
from pdf_scanner_server import output_pdf
#arr = os.listdir('./pdfs_fot_test')

pdffiles = []
base_path_output = "./pdf_outputs/"
i = 0
for file in glob.glob("./pdfs_fot_test/*.pdf"):
    pdffiles.append(file)
    o_path = base_path_output + str(i) + ".txt"
    with codecs.open(o_path, "w", "utf-8") as o:
        o.write(output_pdf(file, 'heb'))
    i += 1
    print("parsed file: " + file)


print(pdffiles)



import os
import subprocess

# start the service
dirname = os.path.dirname(__file__)
# filename = os.path.join(dirname,'../pdf_parser_service/pdf_parser_service.exe')
filename = os.path.join(dirname,'dist/pdf_parser_service/pdf_parser_service.exe')
os.startfile(filename)

# filename_client = os.path.join(dirname,'../pdf_parser_client/pdf_parser_client.exe')
filename_client = os.path.join(dirname,'dist/pdf_parser_client/pdf_parser_client.exe')
subprocess.call(filename_client)
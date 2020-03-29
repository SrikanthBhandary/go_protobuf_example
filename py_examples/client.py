#! /usr/bin/env python

# See README.txt for information and build instructions.

from __future__ import print_function

import addressbook_pb2
import requests


def makerequest(data):
    response = requests.post(
        "http://127.0.0.1:8888/people",
        headers={'Content-Type': 'application/octet-stream'},
        data=data)
    print(response)


if __name__ == "__main__":
    address_book = addressbook_pb2.AddressBook()
    with open("../addressbook.dat", "rb") as f:
        addresss = address_book.ParseFromString(f.read())
        if len(address_book.people) > 0:
            person = address_book.people[0]
            binary = address_book.SerializeToString()
            makerequest(binary)

#!/bin/bash
rm -rf config/*

mv web/mirrors.html .
mv web/admin.html .
rm -rf web/*
mv *.html web/

rm -rf .gitignore

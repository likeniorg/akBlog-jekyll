#!/bin/bash
rm -rf config/*

mv web/mirrors.html .
rm -rf web/*
mv mirrors.html web/

rm -rf .gitignore

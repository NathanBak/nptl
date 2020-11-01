# Nathan's Pseudo Translation Library

## Intro
Nathan's Pseudo Translation Library (nptl) is a golang library that can be customized and used to perform different types of pseudo translations on various formats.

## Basic Architecture
Key concepts in nptl include:
- Runes - instead of byte arrays (which are encoding-specific) and strings (which are byte arrays underneath and also may perform some unwanted normalization), Runes are used to pass around translatable content.
- Translator - a translator is able to take Runes and pseudo translate them into another set of Runes.
- Handler - a handler is able to find the translatable content with a particular format and pass the content to a Translator.
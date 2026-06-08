# Wordle ML - Wordlists

Part of the Wordle ML project. These wordlists represent 5-letter words used by Wordle and Wordle ML models.

## Wordlists

- `data/wordlist-valid-guesses.csv`: words that you are allowed to play in a Wordle game.
- `data/wordlist-valid-solutions.csv`: words that are allowed to be the correct answer.
- `data/wordlist-action-space.csv`: recommended action space for machine learning models.

## Number of Words

- 12,947 valid guesses
- 2,309 valid solutions
- 4,739 action space words

The solutions are a subset of the guesses. The action space list includes all valid solutions plus additional recommended
model actions.

## Format

The files contain one word per line. Words should be in block capitals throughout the project.

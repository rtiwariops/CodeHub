import spacy

nlp = spacy.load("en_core_web_sm")

def extract_entities(text):
    doc = nlp(text)
    entities = [(ent.text, ent.label_) for ent in doc.ents]
    return entities

text = "The patient visited the clinic on February 12, 2023, and received a flu shot."
entities = extract_entities(text)
print(entities)

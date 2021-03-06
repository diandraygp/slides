from Bio import Entrez, SeqIO

Entrez.email = "gabor@szabgab.com"

#doc_id = 'MK792700.1'
doc_id = "EU490707"

# rettype="fasta"
handle = Entrez.efetch(db="nucleotide", id=doc_id, rettype="gb", retmode="text")
data = handle.read()
#print(data)

filename = "temp.data"
with open(filename, 'w') as fh:
    fh.write(data)

file_type = "genbank"
for seq_record in SeqIO.parse(filename, file_type):
    print(seq_record.id)
    print(repr(seq_record.seq))
    print(len(seq_record))

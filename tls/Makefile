default:
.PHONY: default

genca:
	@cfssl gencert \
		-initca ca-csr.json | cfssljson -bare ca
.PHONY: genca

gencert:
	@cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=ca-config.json \
		-profile=server \
		web-csr.json | cfssljson -bare web
.PHONY: gencert

copy:
	# web certs
	@cp ca.pem ../etc/certs
	@cp web.pem ../etc/certs
	@cp web-key.pem ../etc/certs
.PHONY: copy

clear:
	@rm *.{pem,csr}
.PHONY: clear

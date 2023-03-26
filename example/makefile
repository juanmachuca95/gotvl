# Generate translations (en, es)
# Create by definitions
.PHONY: init
init:
	mkdir translations && cd translations; touch active.en.toml active.es.toml

.PHONY: gen
gen:
	cd translations && goi18n merge active.en.toml active.es.toml 

# Use the Finish command only when all translations have been completed.
.PHONY: finish
finish:
	cd translations; echo "\n" >> active.es.toml; cat translate.es.toml >> active.es.toml;

.PHONY: reset
reset: 
	cd translations; rm -rf active.es.toml translate.es.toml; touch active.en.toml

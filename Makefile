.PHONY: build
build:
	@mkdir -p bin
	@cd step-$(DOJO_FX_STEP) && go build -o ../bin/step-$(DOJO_FX_STEP)

.PHONY: run-step
run-step: build
	@./bin/step-$(DOJO_FX_STEP)

.PHONY: watch-step
watch-step:
	@find ./step-$(DOJO_FX_STEP) -name "*.go" | entr -r make run-$(DOJO_FX_STEP)  --no-print-directory

.PHONY: run-00
run-00: DOJO_FX_STEP=00
run-00: run-step

.PHONY: watch-00
watch-00: DOJO_FX_STEP=00
watch-00: watch-step

.PHONY: run-01
run-01: DOJO_FX_STEP=01
run-01: run-step

.PHONY: watch-01
watch-01: DOJO_FX_STEP=01
watch-01: watch-step

.PHONY: run-02
run-02: DOJO_FX_STEP=02
run-02: run-step

.PHONY: watch-02
watch-02: DOJO_FX_STEP=02
watch-02: watch-step

.PHONY: run-03
run-03: DOJO_FX_STEP=03
run-03: run-step

.PHONY: watch-03
watch-03: DOJO_FX_STEP=03
watch-03: watch-step

.PHONY: run-04
run-04: DOJO_FX_STEP=04
run-04: run-step

.PHONY: watch-04
watch-04: DOJO_FX_STEP=04
watch-04: watch-step

.PHONY: run-05
run-05: DOJO_FX_STEP=05
run-05: run-step

.PHONY: watch-05
watch-05: DOJO_FX_STEP=05
watch-05: watch-step

.PHONY: run-06
run-06: DOJO_FX_STEP=06
run-06: run-step

.PHONY: watch-06
watch-06: DOJO_FX_STEP=06
watch-06: watch-step

.PHONY: run-07
run-07: DOJO_FX_STEP=07
run-07: run-step

.PHONY: watch-07
watch-07: DOJO_FX_STEP=07
watch-07: watch-step

.PHONY: run-08
run-08: DOJO_FX_STEP=08
run-08: run-step

.PHONY: watch-08
watch-08: DOJO_FX_STEP=08
watch-08: watch-step

.PHONY: run-09
run-09: DOJO_FX_STEP=09
run-09: run-step

.PHONY: watch-09
watch-09: DOJO_FX_STEP=09
watch-09: watch-step

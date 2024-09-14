document.addEventListener("alpine:init", () => {
    Alpine.data("select1", () => ({
        data: [
            { label: "Foo", value: 1 },
            { label: "Bar", value: 2 },
            { label: "Baz", value: 3 }
        ],
        defaultValue: null,
        get optionLabel() {
            return this.defaultValue
                ? this.data.find((item) => item.value === this.defaultValue).label
                : "-";
        }
    }));
    Alpine.data("select2", () => ({
        defaultValue: null
    }));
    Alpine.data("select3", () => ({
        data: [
            { label: "Foo", value: 1 },
            { label: "Bar", value: 2 },
            { label: "Baz", value: 3 }
        ],
        defaultValue: 3,
        get optionLabel() {
            return this.defaultValue
                ? this.data.find((item) => item.value === this.defaultValue).label
                : "-";
        }
    }));

    Alpine.data("select", (options, value, placeholder) => ({
        options: options,
        placeholder,
        status: false,
        activeIndex: 0,
        selectedIndex: null,
        selectedItem: null,
        init() {
            if (value) {
                this.selectedIndex = this.options.findIndex(
                    (item) => item.value === value
                );
                this.selectedItem = this.options.find((item) => item.value === value);
            }
        },
        onClick() {
            if (this.status) {
                return;
            }

            this.activeIndex = this.selectedIndex ?? this.activeIndex ?? 0;
            this.status = true;
            this.$nextTick(() => {
                setTimeout(() => this.$refs.listbox.focus(), 10);
            });
        },

        selectItem(item, index) {
            this.selectedIndex = index;
            this.selectedItem = item;
            this.status = false;

            this.$dispatch("input", item.value);
        },

        button: {
            ["x-ref"]: "button",
            ["@click"]() {
                this.onClick();
            },
            ["@keydown.arrow-up.stop.prevent"]() {
                this.onClick();
            },
            ["@keydown.arrow-down.stop.prevent"]() {
                this.onClick();
            }
        },

        listbox: {
            ["x-transition:leave"]: "transition ease-in duration-100",
            ["x-transition:leave-start"]: "opacity-100",
            ["x-transition:leave-end"]: "opacity-0",
            ["x-ref"]: "listbox",
            ["x-show"]() {
                return this.status;
            },
            ["@click.outside"]() {
                this.status = false;
            },
            ["@keydown.enter.stop.prevent"]() {
                if (this.activeIndex === null) {
                    return;
                }

                this.selectItem(this.options[this.activeIndex], this.activeIndex);
                this.$refs.button.focus();
            },
            ["@keydown.space.stop.prevent"]() {
                if (this.activeIndex === null) {
                    return;
                }

                this.selectItem(this.options[this.activeIndex], this.activeIndex);
                this.$refs.button.focus();
            },
            ["@keydown.escape"]() {
                this.status = false;
                this.$refs.button.focus();
            },
            ["@keyup.arrow-up"]() {
                this.activeIndex =
                    this.activeIndex - 1 < 0
                        ? this.options.length - 1
                        : this.activeIndex - 1;
            },
            ["@keyup.arrow-down"]() {
                this.activeIndex =
                    this.activeIndex + 1 > this.options.length - 1
                        ? 0
                        : this.activeIndex + 1;
            }
        },

        list: {
            ["x-for"]: "(template, index) in options"
        },

        listItem: {
            [":id"]() {
                return `option-${this.index}`;
            },
            [":class"]() {
                const isActive = this.activeIndex === this.index;
                return {
                    "text-white bg-indigo-600": isActive,
                    "text-gray-900": !isActive
                };
            },
            ["@click"]() {
                this.selectItem(this.template, this.index);
            },
            ["@mouseleave"]() {
                this.activeIndex = null;
            },
            ["@mouseenter"]() {
                this.activeIndex = this.index;
            }
        },

        listItemLabel: {
            ["x-text"]() {
                return this.template.label;
            },
            [":class"]() {
                const isSelected = this.selectedIndex === this.index;
                return { "font-semibold": isSelected, "font-normal": !isSelected };
            }
        },
        listItemCheckIcon: {
            ["x-show"]() {
                return this.selectedIndex === this.index;
            },
            [":class"]() {
                const isActive = this.activeIndex === this.index;
                return { "text-white": isActive, "text-indigo-600": !isActive };
            }
        },
        selectedItemLabel: {
            ["x-text"]() {
                return this.selectedItem?.label || this.placeholder;
            }
        },
        caretIcon: {
            [":class"]() {
                return { "fa-angle-up": this.status, "fa-angle-down": !this.status };
            }
        }
    }));
});

<template>
    <label for="toggle" id="tb">
        <input type="checkbox" id="toggle" :checked="checked" @click.stop="toggle">
        <span></span>
        <div v-if="checked" id="on">Mine</div>
        <div v-if="!checked" id="off">All</div>
    </label>
</template>

<script>
export default {
    props: {
        value: {
            type: Boolean,
            default: false
        },
    },
    data() {
        return {
            checked: this.value,
        }
    },
    watch: {
        value() {
            this.checked = this.value;
        }
    },
    methods: {
        toggle() {
            setTimeout(() => {
                this.checked = !this.checked
                this.$emit('toggle', this.checked);
            }, 100);
            // console.log(this.value)
        }
    },
    components: {}
}
</script>

<style scoped>
span,
#open,
#tb,
#off {
    --button-width: 60px;
    --button-height: 29.5px;
    --toggle-diameter: 25.5px;
    --button-toggle-offset: calc((var(--button-height) - var(--toggle-diameter)) / 2);
    --toggle-shadow-offset: 10px;
    --toggle-wider: 33px;
    --color-grey: #E9E9E9;
    --color-dark-grey: #39393D;
    --color-green: #30D158;
}

#tb {
    position: relative;
    /* font-weight: bold; */
    cursor: pointer;
    width: var(--button-width);
    height: var(--button-height);
}


#on {
    position: absolute;
    left: 8px;
    top: 0;
    line-height: calc(var(--button-height) + 2px);
    font-size: 10px;
    color: #fff;
    z-index: 0;
}

#off {
    position: absolute;
    left: calc(var(--button-width) / 1.7);
    top: 0;
    line-height: calc(var(--button-height) + 2px);
    font-size: 10px;
    color: gray;
    z-index: 0;
}

#toggle {
    z-index: 1;
}

span {
    display: inline-block;
    width: var(--button-width);
    height: var(--button-height);
    background-color: var(--color-grey);
    border-radius: calc(var(--button-height) / 2);
    position: absolute;
    transition: .2s all ease-in-out;
}

span::after {
    content: '';
    display: inline-block;
    width: var(--toggle-diameter);
    height: var(--toggle-diameter);
    background-color: #fff;
    border-radius: calc(var(--toggle-diameter) / 2);
    position: absolute;
    top: var(--button-toggle-offset);
    transform: translateX(var(--button-toggle-offset));
    box-shadow: var(--toggle-shadow-offset) 0 calc(var(--toggle-shadow-offset) * 4) rgba(0, 0, 0, .10);
    transition: .2s all ease-in-out;
}

input[type="checkbox"]:checked + span {
    background-color: var(--color-green);
}

input[type="checkbox"]:checked + span::after {
    transform: translateX(calc(var(--button-width) - var(--toggle-diameter) - var(--button-toggle-offset)));
    box-shadow: calc(var(--toggle-shadow-offset) * -1) 0 calc(var(--toggle-shadow-offset) * 4) rgba(0, 0, 0, .10);
}

input[type="checkbox"] {
    display: none;
}

input[type="checkbox"]:active + span::after {
    width: var(--toggle-wider);
}

input[type="checkbox"]:checked:active + span::after {
    transform: translateX(calc(var(--button-width) - var(--toggle-wider) - var(--button-toggle-offset)));
}

</style>
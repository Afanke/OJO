<template>
    <div class="code-input-main">
        <div
            class="code-input-main-item"
            v-for="(item,index) in codeList"
            :key="index">
            {{ code[index] || '' }}
        </div>
        <input class="code-input-input" v-model="code" maxlength="6" type="number"/>
    </div>
</template>

<script>
export default {
    name: "CodeInput",
    props: ['value'],
    data() {
        return {
            codeList: [],
            code: "",
            codeLength: 6,
        };
    },
    mounted() {
        // 定义一个数组用于循环
        this.codeList = new Array(this.codeLength).fill("");
    },
    watch: {
        // 截取字符长度
        code() {
            if (this.code.length > this.codeLength) {
                this.code = this.code.substring(0, this.codeLength);
            }
            console.log(this.code)
            this.$emit('input', this.code)
        }
    },
    methods: {
        getCode() {
            return this.code;
        }
    }
};
</script>
<style scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
    -webkit-appearance: none !important;
    margin: 0;
}

.code-input-main {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    position: relative;
}

.code-input-input {
    height: 44px;
    width: 100%;
    position: absolute;
    border: none;
    outline: none;
    color: transparent;
    background-color: transparent;
    text-shadow: 0 0 0 transparent;
}

.code-input-main-item {
    width: 34px;
    height: 44px;
    margin: 0 5px;
    padding-bottom: 0;
    opacity: 0.8;
    border-bottom: solid #323232 1px;
    text-align: center;
    font-size: 30px;
    color: #323232;
}
</style>
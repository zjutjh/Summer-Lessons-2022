<template>
  <div class="child">
    <h2>hello world</h2>
    <div>这是一个组件模板</div>
    <div>{{ msg }}</div>
    <button @click="handleClick">点击修改父元素标题</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "HelloWorld",
  components: {},
  props: {
    msg: String,
    titleCopy: String,
  },
  data() {},
  methods: {
    handleClick() {
      this.$emit("update:titleCopy", "father");
    },
  },
  computed: {},
  mounted() {},
});
</script>

<style lang="scss">
.child {
  margin-top: 150px;
  border: 1px solid black;
}
</style>

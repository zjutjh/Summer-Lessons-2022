<template>
  <h1>{{ title }}</h1>
  <div>这是一个组件模板</div>
  <hello-word :msg="message" v-model:titleCopy="title"></hello-word>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HelloWord from "./components/HelloWorld.vue";

export default defineComponent({
  components: { HelloWord },
  props: {},
  data() {
    return {
      message: "hi",
      title: "App.vue",
    };
  },
  methods: {},
  computed: {},
  mounted() {},
});
</script>

<style lang="scss"></style>

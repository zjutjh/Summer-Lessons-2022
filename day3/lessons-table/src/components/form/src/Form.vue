<template>
  <form class="k-form">
    <slot></slot>
  </form>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "Form",
  setup(props, context) {},
});
</script>

<style lang="scss" scoped>
.k-form {
  width: 80%;
  display: flex;
  flex-direction: column;
  row-gap: 10px;
}
</style>

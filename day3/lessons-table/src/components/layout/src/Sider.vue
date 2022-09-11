<template>
  <aside class="k-sider">
    <slot></slot>
  </aside>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {},
});
</script>

<style lang="scss">
.k-sider {
  background: #fff;
  width: 250px;
  border: 1px solid rgb(0, 0, 0, 0.06);
  display: flex;
  flex-direction: column;
}

.k-sider > div {
  border-bottom: 1px solid rgb(0, 0, 0, 0.06);
  padding: 20px;
  cursor: pointer;
}
.k-sider > div:hover {
  background-color: #fafafa;
}
</style>

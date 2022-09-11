<template>
  <div class="k-card">
    <div v-if="title" class="k-card-head">
      <h1 class="k-card-title">{{ title }}</h1>
    </div>
    <div class="k-card-body">
      <slot></slot>
    </div>
    <div v-if="hasFooter" class="k-card-footer">
      <slot name="footer"></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Card",
  props: {
    title: String,
  },
  setup(props, context) {
    const hasFooter = !!context.slots.footer;
    return {
      hasFooter,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/styles";
.k-card {
  box-sizing: border-box;
  border: 1px solid rgb(0, 0, 0, 0.06);
  background: #fff;
}
.k-card-head {
  padding: 0 24px;
  border-bottom: 1px solid rgb(0, 0, 0, 0.06);
  .k-card-title {
    font-size: 20px;
  }
}
.k-card-body {
  padding: 24px;
}
.k-card-footer {
  border-top: 1px solid rgb(0, 0, 0, 0.06);
  padding: 10px 24px;
  text-align: end;
}
</style>

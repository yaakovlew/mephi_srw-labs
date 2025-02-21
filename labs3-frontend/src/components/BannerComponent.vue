<template>
  <div class="banner-component q-px-lg q-py-sm flex column g-m">
    <div class="banner-title text-primary">
      <slot name="title" />
    </div>
    <div class="banner-content">
      <slot />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.banner-component {
  background-color: white;
  border: 1px solid white;
  border-radius: 5px;
}

.banner-title {
  font-weight: 600;
  font-size: 18px;
}

.banner-content {
  width: 100%;
}
</style>

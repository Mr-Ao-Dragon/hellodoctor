<script setup lang="ts">

import {ref} from 'vue';
import {Icon as TIcon} from 'tdesign-icons-vue-next';

const accessToken = localStorage.getItem('access_token');
const expiresIn = localStorage.getItem('expires_in');

if (!accessToken || !expiresIn || parseInt(expiresIn, 10) <= Math.floor(Date.now() / 1000)) {
  window.location.href = '/entrypoint';
}

const value = ref('label_1');
const list = ref([
  {name: 'label_1', text: '首页', icon: 'home', badgeProps: {}, ariaLabel: '首页'},
  {name: 'label_2', text: '聊天', icon: 'chat', badgeProps: {}, ariaLabel: '聊天'},
  {name: 'label_3', text: '我的', icon: 'user', badgeProps: {}, ariaLabel: '我的'},
]);
</script>
<template>
  <t-tab-bar v-model="value" :split="false">
    <t-tab-bar-item
        v-for="item in list"
        :key="item.name"
        :value="item.name"
        :badge-props="item.badgeProps"
        :aria-label="item.ariaLabel"
    >
      {{ item.text }}
      <template #icon>
        <t-icon :name="item.icon"/>
      </template>
    </t-tab-bar-item>
  </t-tab-bar>
</template>


<style scoped>
</style>

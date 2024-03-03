<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition
        mode="out-in"
        name="el-fade-in-linear"
      >
        <keep-alive :include="routerStore.keepAliveRouters">
          <div>
            <component :is="Component" />
          </div>
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
import { useRouterStore } from '@/pinia/modules/router'
const routerStore = useRouterStore()
defineOptions({
  name: 'Biz'
})
</script>

<template>
  <div>
    <el-table
      ref="multipleTable"
      :data="tableData"
      style="width: 100%"
      tooltip-effect="dark"
      row-key="ID"
    >
      <el-table-column
        align="left"
        label="姓名"
        prop="collaboratorName"
        width="120"
      />
      <el-table-column
        align="left"
        label="手机号"
        prop="phoneNumber"
        width="120"
      />
      <el-table-column
        align="left"
        label="操作"
        min-width="160"
      />
    </el-table>
  </div>
</template>

<script setup>
import { ref } from 'vue'

defineOptions({
  name: 'Members'
})

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const tableData = ref([])

// const init = () => {
//   props.row.dataAuthorityId && props.row.dataAuthorityId.forEach(item => {
//     const obj = authoritys.value && authoritys.value.filter(au => au.authorityId === item.authorityId) && authoritys.value.filter(au => au.authorityId === item.authorityId)[0]
//     dataAuthorityId.value.push(obj)
//   })
// }
//
// init()
//
// // 暴露给外层使用的切换拦截统一方法
// const enterAndNext = () => {
//   authDataEnter()
// }
//
// defineExpose({
//   enterAndNext
// })

</script>

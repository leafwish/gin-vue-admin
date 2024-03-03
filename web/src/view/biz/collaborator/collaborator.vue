<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="机构名称"
          prop="teamName"
          width="120"
        />
        <el-table-column
          align="left"
          label="负责人"
          prop="leaderName"
          width="120"
        />
        <el-table-column
          align="left"
          label="联系电话"
          prop="leaderPhoneNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="创建时间"
          width="180"
        >
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="更新时间"
          width="180"
        >
          <template #default="scope">
            <span>{{ formatDate(scope.row.UpdatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="160"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editCollaboratorTeam(scope.row)"
            >变更</el-button>
            <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >取消</el-button>
                <el-button
                  type="primary"
                  @click="removeCollaboratorTeam(scope.row)"
                >确定</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                  @click="scope.row.visible = true"
                >删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="机构"
    >
      <el-form
        :inline="true"
        :model="form"
        label-width="80px"
      >
        <el-form-item label="机构名称">
          <el-input
            v-model="form.teamName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="负责人">
          <el-input
            v-model="form.leaderName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input
            v-model="form.leaderPhoneNumber"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCollaboratorTeam,
  updateCollaboratorTeam,
  deleteCollaboratorTeam,
  getCollaboratorTeam,
  getCollaboratorTeamPage,
} from '@/api/biz/collaborator'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'Collaborator'
})

const form = ref({
  teamName: '',
  leaderName: '',
  leaderPhoneNumber: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCollaboratorTeamPage({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const dialogFormVisible = ref(false)
const type = ref('')
const editCollaboratorTeam = async(row) => {
  const res = await getCollaboratorTeam({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    form.value = res.data.collaboratorTeam
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    teamName: '',
    leaderName: '',
    leaderPhoneNumber: ''
  }
}
const removeCollaboratorTeam = async(row) => {
  row.visible = false
  const res = await deleteCollaboratorTeam({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createCollaboratorTeam(form.value)
      break
    case 'update':
      res = await updateCollaboratorTeam(form.value)
      break
    default:
      res = await createCollaboratorTeam(form.value)
      break
  }

  if (res.code === 0) {
    closeDialog()
    getTableData()
  }
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

</script>

<style></style>

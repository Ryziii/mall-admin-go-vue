<template>
  <div class="user-list">
    <h1>权限管理</h1>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="username"
        label="用户名">
      </el-table-column>
      <el-table-column
        prop="password"
        label="密码">
      </el-table-column>
      <el-table-column
        prop="role_name"
        label="权限">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-button @click="editRow(scope.row)" size="small">编辑</el-button>
          <el-button @click="handleDelete(scope.row)" size="small" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button class="add-user" type="primary" @click.native="dialogFormVisible=true">新增用户</el-button>

    <el-dialog :title="dialogType === 'edit' ? '编辑用户' : '新增用户'" :visible.sync="dialogFormVisible" :close-on-click-modal="false" destroy-on-close ref="dialogForm" @close="closeDialog">
      <el-form :model="form" :rules="rules" ref="ruleForm">
        <el-form-item label="用户角色" label-width="120px" prop="role_id">
          <el-select v-model="form.role_id" placeholder="请选择用户角色">
            <el-option
              v-for="item in roles"
              :key="item.role_id"
              :label="item.role_name"
              :value="item.role_id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="用户名" label-width="120px" prop="username">
          <el-input v-model="form.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="120px" prop="password">
          <el-input v-model="form.password" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitForm('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { cloneDeep } from "lodash"
import PermissionService from "@/service/permission.service.js"

export default {
  name: 'category-list',
  data() {
    return {
      tableData: [],
      dialogFormVisible: false,
      form: {
        role_name: '',
        username: '',
        password: '',
        role_id: ''
      },
      roles: [],
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'change' }
        ]
      },
      dialogType: ''
    }
  },
  created() {
    this.getPermissionList()
    this.getRoles()
  },
  methods: {
    getPermissionList() {
      PermissionService.getPermissionList().then(res => {
        const { data } = res
        this.tableData = data
      })
    },
    getRoles() {
      PermissionService.getRoles().then(res => {
        const { data } = res
        this.roles = data
      })
    },
    handleDelete(row) {
      this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        PermissionService.deleteAdmin(row.admin_id).then(() => {
          this.getPermissionList()
          this.$message({
            type: 'success',
            message: '用户删除成功!'
          });
        })
      });
      
    },
    editRow(row) {
      this.form = cloneDeep(row);
      this.dialogType = 'edit';
      this.dialogFormVisible = true;
    },
    closeDialog() {
      this.form = {};
      this.dialogFormVisible = false;
      this.dialogType = '';
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          PermissionService.updateRole(this.form.admin_id, this.form).then(() => {
            this.getPermissionList()
          })
          this.dialogFormVisible = false
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
  }
}
</script>

<style scoped lang="scss">
.user-list {
  padding: 0 20px;
  .add-user {
    float: left;
    margin-top: 20px;
  }
  .el-select {
    width: 100%;
  }
}
</style>
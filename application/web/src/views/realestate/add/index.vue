<template>
  <div class="app-container">
    <el-form
      ref="ruleForm"
      v-loading="loading"
      :model="ruleForm"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item label="owner" prop="proprietor">
        <el-select
          v-model="ruleForm.proprietor"
          placeholder="Please choose the owner"
          @change="selectGet"
        >
          <el-option
            v-for="item in accountList"
            :key="item.accountId"
            :label="item.userName"
            :value="item.accountId"
          >
            <span style="float: left">{{ item.userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{
              item.accountId
            }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Overall space ㎡" prop="totalArea">
        <el-input-number
          v-model="ruleForm.totalArea"
          :precision="2"
          :step="0.1"
          :min="0"
        />
      </el-form-item>
      <el-form-item label="Living Space ㎡" prop="livingSpace">
        <el-input-number
          v-model="ruleForm.livingSpace"
          :precision="2"
          :step="0.1"
          :min="0"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')"
          >Create immediately</el-button
        >
        <el-button @click="resetForm('ruleForm')">Repossess</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { queryAccountList } from "@/api/account";
import { createRealEstate } from "@/api/realEstate";

export default {
  name: "AddRealeState",
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error("Must be greater than 0"));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        proprietor: "",
        totalArea: 0,
        livingSpace: 0,
      },
      accountList: [],
      rules: {
        proprietor: [
          {
            required: true,
            message: "Please choose the owner",
            trigger: "change",
          },
        ],
        totalArea: [{ validator: checkArea, trigger: "blur" }],
        livingSpace: [{ validator: checkArea, trigger: "blur" }],
      },
      loading: false,
    };
  },
  computed: {
    ...mapGetters(["accountId"]),
  },
  created() {
    queryAccountList().then((response) => {
      if (response !== null) {
        // Filter off the administrator
        this.accountList = response.filter(
          (item) => item.userName !== "administrator"
        );
      }
    });
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm("Whether to create immediately?", "hint", {
            confirmButtonText: "Sure",
            cancelButtonText: "Cancel",
            type: "success",
          })
            .then(() => {
              this.loading = true;
              createRealEstate({
                accountId: this.accountId,
                proprietor: this.ruleForm.proprietor,
                totalArea: this.ruleForm.totalArea,
                livingSpace: this.ruleForm.livingSpace,
              })
                .then((response) => {
                  this.loading = false;
                  if (response !== null) {
                    this.$message({
                      type: "success",
                      message: "Successful creation!",
                    });
                  } else {
                    this.$message({
                      type: "error",
                      message: "Failed to create!",
                    });
                  }
                })
                .catch((_) => {
                  this.loading = false;
                });
            })
            .catch(() => {
              this.loading = false;
              this.$message({
                type: "info",
                message: "Cancel",
              });
            });
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    selectGet(accountId) {
      this.ruleForm.proprietor = accountId;
    },
  },
};
</script>

<style scoped></style>

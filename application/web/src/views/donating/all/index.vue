<template>
  <div class="container">
    <el-alert type="success">
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>Balance: $ {{ balance }} dollar</p>
    </el-alert>
    <div v-if="donatingList.length == 0" style="text-align: center">
      <el-alert title="Can't check the data" type="warning" />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col
        v-for="(val, index) in donatingList"
        :key="index"
        :span="6"
        :offset="1"
      >
        <el-card class="d-all-card">
          <div slot="header" class="clearfix">
            <span>{{ val.donatingStatus }}</span>
            <el-button
              v-if="
                roles[0] !== 'admin' &&
                val.grantee === accountId &&
                val.donatingStatus === 'Donation'
              "
              style="float: right; padding: 3px 6px"
              type="text"
              @click="updateDonating(val, 'done')"
              >Confirm the receiving</el-button
            >
            <el-button
              v-if="
                roles[0] !== 'admin' &&
                (val.donor === accountId || val.grantee === accountId) &&
                val.donatingStatus === 'Donation'
              "
              style="float: right; padding: 3px 0"
              type="text"
              @click="updateDonating(val, 'cancelled')"
              >Cancel</el-button
            >
          </div>
          <div class="item">
            <el-tag>Real estate ID: </el-tag>
            <span>{{ val.objectOfDonating }}</span>
          </div>
          <div class="item">
            <el-tag type="success">Donor ID: </el-tag>
            <span>{{ val.donor }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">Donor ID: </el-tag>
            <span>{{ val.grantee }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">Creation time: </el-tag>
            <span>{{ val.createTime }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { queryDonatingList, updateDonating } from "@/api/donating";

export default {
  name: "AllDonating",
  data() {
    return {
      loading: true,
      donatingList: [],
    };
  },
  computed: {
    ...mapGetters(["accountId", "roles", "userName", "balance"]),
  },
  created() {
    queryDonatingList()
      .then((response) => {
        if (response !== null) {
          this.donatingList = response;
        }
        this.loading = false;
      })
      .catch((_) => {
        this.loading = false;
      });
  },
  methods: {
    updateDonating(item, type) {
      let tip = "";
      if (type === "done") {
        tip = "Confirm that accepting donations";
      } else {
        tip = "Cancel donation operation";
      }
      this.$confirm("Whether" + tip + "?", "hint", {
        confirmButtonText: "Sure",
        cancelButtonText: "Cancel",
        type: "success",
      })
        .then(() => {
          this.loading = true;
          updateDonating({
            donor: item.donor,
            grantee: item.grantee,
            objectOfDonating: item.objectOfDonating,
            status: type,
          })
            .then((response) => {
              this.loading = false;
              if (response !== null) {
                this.$message({
                  type: "success",
                  message: tip + "Successful operation!",
                });
              } else {
                this.$message({
                  type: "error",
                  message: tip + "operation failed!",
                });
              }
              setTimeout(() => {
                window.location.reload();
              }, 1000);
            })
            .catch((_) => {
              this.loading = false;
            });
        })
        .catch(() => {
          this.loading = false;
          this.$message({
            type: "info",
            message: "Cancelled" + tip,
          });
        });
    },
  },
};
</script>

<style>
.container {
  width: 100%;
  text-align: center;
  min-height: 100%;
  overflow: hidden;
}
.tag {
  float: left;
}

.item {
  font-size: 14px;
  margin-bottom: 18px;
  color: #999;
}

.clearfix:before,
.clearfix:after {
  display: table;
}
.clearfix:after {
  clear: both;
}

.d-all-card {
  width: 280px;
  height: 300px;
  margin: 18px;
}
</style>

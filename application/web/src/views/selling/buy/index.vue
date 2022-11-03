<template>
  <div class="container">
    <el-alert type="success">
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>Balance: ￥{{ balance }} 元</p>
    </el-alert>
    <div v-if="sellingList.length == 0" style="text-align: center">
      <el-alert title="Can't check the data" type="warning" />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col
        v-for="(val, index) in sellingList"
        :key="index"
        :span="6"
        :offset="1"
      >
        <el-card class="buy-card">
          <div slot="header" class="clearfix">
            <span>{{ val.selling.sellingStatus }}</span>
            <el-button
              v-if="
                val.selling.sellingStatus !== 'Finish' &&
                val.selling.sellingStatus !== 'expired' &&
                val.selling.sellingStatus !== 'Cancelled'
              "
              style="float: right; padding: 3px 0"
              type="text"
              @click="updateSelling(val, 'cancelled')"
              >取消</el-button
            >
          </div>
          <div class="item">
            <el-tag type="warning">order time: </el-tag>
            <span>{{ val.createTime }}</span>
          </div>
          <div class="item">
            <el-tag>Real estate ID: </el-tag>
            <span>{{ val.selling.objectOfSale }}</span>
          </div>
          <div class="item">
            <el-tag type="success">Seller ID: </el-tag>
            <span>{{ val.selling.seller }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">price: </el-tag>
            <span>￥{{ val.selling.price }} 元</span>
          </div>
          <div class="item">
            <el-tag type="warning">Validity period: </el-tag>
            <span>{{ val.selling.salePeriod }} 天</span>
          </div>
          <div class="item">
            <el-tag type="info">Creation time: </el-tag>
            <span>{{ val.selling.createTime }}</span>
          </div>
          <div class="item">
            <el-tag>Buyer ID: </el-tag>
            <span v-if="val.selling.buyer === ''">Wait</span>
            <span>{{ val.selling.buyer }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { querySellingListByBuyer, updateSelling } from "@/api/selling";

export default {
  name: "BuySelling",
  data() {
    return {
      loading: true,
      sellingList: [],
    };
  },
  computed: {
    ...mapGetters(["accountId", "userName", "balance"]),
  },
  created() {
    querySellingListByBuyer({ buyer: this.accountId })
      .then((response) => {
        if (response !== null) {
          this.sellingList = response;
        }
        this.loading = false;
      })
      .catch((_) => {
        this.loading = false;
      });
  },
  methods: {
    updateSelling(item, type) {
      let tip = "";
      if (type === "done") {
        tip = "confirmed paid";
      } else {
        tip = "Cancel operation";
      }
      this.$confirm("Whether" + tip + "?", "hint", {
        confirmButtonText: "Sure",
        cancelButtonText: "Cancel",
        type: "success",
      })
        .then(() => {
          this.loading = true;
          updateSelling({
            buyer: item.selling.buyer,
            objectOfSale: item.selling.objectOfSale,
            seller: item.selling.seller,
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

.buy-card {
  width: 280px;
  height: 430px;
  margin: 18px;
}
</style>

(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5c0e61bc"],{"00f7":function(e,t,r){"use strict";r.r(t);var o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("el-form",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],ref:"ruleForm",attrs:{model:e.ruleForm,rules:e.rules,"label-width":"100px"}},[r("el-form-item",{attrs:{label:"owner",prop:"proprietor"}},[r("el-select",{attrs:{placeholder:"Please choose the owner"},on:{change:e.selectGet},model:{value:e.ruleForm.proprietor,callback:function(t){e.$set(e.ruleForm,"proprietor",t)},expression:"ruleForm.proprietor"}},e._l(e.accountList,(function(t){return r("el-option",{key:t.accountId,attrs:{label:t.userName,value:t.accountId}},[r("span",{staticStyle:{float:"left"}},[e._v(e._s(t.userName))]),r("span",{staticStyle:{float:"right",color:"#8492a6","font-size":"13px"}},[e._v(e._s(t.accountId))])])})),1)],1),r("el-form-item",{attrs:{label:"Overall space ㎡",prop:"totalArea"}},[r("el-input-number",{attrs:{precision:2,step:.1,min:0},model:{value:e.ruleForm.totalArea,callback:function(t){e.$set(e.ruleForm,"totalArea",t)},expression:"ruleForm.totalArea"}})],1),r("el-form-item",{attrs:{label:"Living Space ㎡",prop:"livingSpace"}},[r("el-input-number",{attrs:{precision:2,step:.1,min:0},model:{value:e.ruleForm.livingSpace,callback:function(t){e.$set(e.ruleForm,"livingSpace",t)},expression:"ruleForm.livingSpace"}})],1),r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.submitForm("ruleForm")}}},[e._v("Create immediately")]),r("el-button",{on:{click:function(t){return e.resetForm("ruleForm")}}},[e._v("Repossess")])],1)],1)],1)},a=[],n=r("5530"),l=(r("d9e2"),r("4de4"),r("d3b7"),r("2f62")),i=r("5723"),c=r("1c0b"),s={name:"AddRealeState",data:function(){var e=function(e,t,r){t<=0?r(new Error("Must be greater than 0")):r()};return{ruleForm:{proprietor:"",totalArea:0,livingSpace:0},accountList:[],rules:{proprietor:[{required:!0,message:"Please choose the owner",trigger:"change"}],totalArea:[{validator:e,trigger:"blur"}],livingSpace:[{validator:e,trigger:"blur"}]},loading:!1}},computed:Object(n["a"])({},Object(l["b"])(["accountId"])),created:function(){var e=this;Object(i["b"])().then((function(t){null!==t&&(e.accountList=t.filter((function(e){return"administrator"!==e.userName})))}))},methods:{submitForm:function(e){var t=this;this.$refs[e].validate((function(e){if(!e)return!1;t.$confirm("Whether to create immediately?","hint",{confirmButtonText:"Sure",cancelButtonText:"Cancel",type:"success"}).then((function(){t.loading=!0,Object(c["a"])({accountId:t.accountId,proprietor:t.ruleForm.proprietor,totalArea:t.ruleForm.totalArea,livingSpace:t.ruleForm.livingSpace}).then((function(e){t.loading=!1,null!==e?t.$message({type:"success",message:"Successful creation!"}):t.$message({type:"error",message:"Failed to create!"})})).catch((function(e){t.loading=!1}))})).catch((function(){t.loading=!1,t.$message({type:"info",message:"Cancel"})}))}))},resetForm:function(e){this.$refs[e].resetFields()},selectGet:function(e){this.ruleForm.proprietor=e}}},u=s,p=r("2877"),m=Object(p["a"])(u,o,a,!1,null,"67278c23",null);t["default"]=m.exports},"1c0b":function(e,t,r){"use strict";r.d(t,"a",(function(){return a})),r.d(t,"b",(function(){return n}));var o=r("b775");function a(e){return Object(o["a"])({url:"/createRealEstate",method:"post",data:e})}function n(e){return Object(o["a"])({url:"/queryRealEstateList",method:"post",data:e})}}}]);
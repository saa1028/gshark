<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/nav_logo.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">GShark</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="loginForm"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                width="100%"
                height="100%"
                alt="请输入验证码"
                @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
/* eslint-disable vue/multi-word-component-names */
import { mapActions } from "vuex";
import { captcha } from "@/api/user";
export default {
  name: "Login",
  data() {
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        username: "",
        password: "",
        captcha: "",
        captchaId: "",
      },
      logVerify: "",
      picPath: "",
    };
  },
  created() {
    this.loginVefify();
    this.curYear = new Date().getFullYear();
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      return await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      // const flag = await this.login();
      // if (!flag) {
      //   this.loginVefify();
      // }
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          const flag = await this.login();
          if (!flag) {
            this.loginVefify();
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loginVefify();
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    async loginVefify() {
      const result = await captcha({});
      this.picPath = result.data.picPath;
      this.loginForm.captchaId = result.data.captchaId;
      // captcha({}).then((ele) => {
      //   this.picPath = ele.data.picPath;
      //   this.loginForm.captchaId = ele.data.captchaId;
      // });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>

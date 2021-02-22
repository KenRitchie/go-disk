<template>


  <el-container>

    <el-header style="font-size: 10px;">
      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="1"><i class="el-icon-folder"></i>所有文件</el-menu-item>

        <el-menu-item index="2"><i class="el-icon-setting"></i>设置</el-menu-item>
      </el-menu>
    </el-header>

    <el-container>

      <el-main>
        xxxx
      </el-main>
    </el-container>
    <el-footer>
      <div class="footer-author">
        <div>
          <i class="el-icon-folder"></i>
          <el-link href="https://github.com/kenritchie/go-disk" target="_blank">
            go-disk
          </el-link>
        </div>
        <div>构建版本 v0.5.4 | Copyright © 2019 KenRitChie</div>

      </div>
    </el-footer>
  </el-container>


</template>

<style>
.footer-author {
  text-align: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>

<script>
export default {
  name: 'FileDetail',
  data() {
    // const item = {
    //   date: '2016-05-02',
    //   name: '王小虎',
    //   size: '1024kb',
    //
    // };
    return {
      // tableData: Array(10).fill(item),
      activeIndex: '1'
    }
  },
  methods: {
    handleSelect(key) {
      console.log(key);
      console.log(typeof (key))
      if (key === "1") {
        this.$router.push({name: 'list'});
      } else {
        this.$router.push({name: 'setting'});
      }

    },
    handleDownload(key){
      console.log(key,)
    }
  }
};
</script>

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
        <el-table :data="tableData"  >

          <el-table-column prop="name" label="名称"    >
            <template slot-scope="scope">
              <router-link :to="{name:'file',param:''}" >{{scope.row.name}}</router-link>
            </template>
          </el-table-column>
          <el-table-column prop="size" label="大小">
          </el-table-column>
          <el-table-column prop="date" label="修改日期">
          </el-table-column>
          <el-table-column prop="operation" fixed="right" width="100" label="操作">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
              <el-button @click="handleClick(scope.row)" type="text" size="small">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
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
  name: 'EContainer',
  data() {
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      size: '1024kb',

    };
    return {
      tableData: Array(10).fill(item),
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
    handleClick(key) {
      console.log(key,)
    }
  }
};
</script>

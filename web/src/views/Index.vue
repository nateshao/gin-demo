<template>
  <el-container>
    <el-header>公众号：程序员千羽</el-header>
    <el-main>
      <el-row type="flex" justify="center">
        <el-col :xs="20" :span="12" >
          <div class="grid-content">
            <el-divider>
              <h1>TODO备忘录</h1>
            </el-divider>
            <TodoList></TodoList>
          </div>
        </el-col>
      </el-row>
    </el-main>
    <el-footer>~~求关注~~</el-footer>
  </el-container>
</template>

<script>
// @ is an alias to /src
import TodoList from "@/components/TodoList.vue";

export default {
  name: "Index",
  components: {
    TodoList
  }
};
</script>
<style>
.el-header,
.el-footer {
  background-color: #409eff;
  color: #fff;
  text-align: center;
  line-height: 60px;
}
.el-footer {
  background-color: #909399;
  display: block;
  width: 100%;
  position: fixed;
  bottom: 0;
}
</style>
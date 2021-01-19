package reflect

import "html/template"

var repositoryTmpl = template.Must(template.New("repository").Parse(`package repositories

import (
	"{{.PkgName}}/model"
	"github.com/enrichroad/community/database"
	"github.com/enrichroad/community/rest"
	"github.com/enrichroad/community/pagination"
	"gorm.io/gorm"
)

var {{.Name}}Repository = new{{.Name}}Repository()

func new{{.Name}}Repository() *{{.CamelName}}Repository {
	return &{{.CamelName}}Repository{}
}

type {{.CamelName}}Repository struct {
}

func (r *{{.CamelName}}Repository) Get(db *gorm.DB, id int64) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) Take(db *gorm.DB, where ...interface{}) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) Find(db *gorm.DB, cnd *database.SqlCnd) (list []model.{{.Name}}) {
	cnd.Find(db, &list)
	return
}

func (r *{{.CamelName}}Repository) FindOne(db *gorm.DB, cnd *database.SqlCnd) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) FindPageByParams(db *gorm.DB, params *rest.QueryParams) (list []model.{{.Name}}, paging *pagination.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *{{.CamelName}}Repository) FindPageByCnd(db *gorm.DB, cnd *database.SqlCnd) (list []model.{{.Name}}, paging *pagination.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.{{.Name}}{})

	paging = &pagination.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *{{.CamelName}}Repository) Count(db *gorm.DB, cnd *database.SqlCnd) int64 {
	return cnd.Count(db, &model.{{.Name}}{})
}

func (r *{{.CamelName}}Repository) Create(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Create(t).Error
	return
}

func (r *{{.CamelName}}Repository) Update(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Save(t).Error
	return
}

func (r *{{.CamelName}}Repository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *{{.CamelName}}Repository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *{{.CamelName}}Repository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.{{.Name}}{}, "id = ?", id)
}

`))

var serviceTmpl = template.Must(template.New("service").Parse(`package services

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/repositories"
	"github.com/enrichroad/community/database"
	"github.com/enrichroad/community/rest"
	"github.com/enrichroad/community/pagination"
)

var {{.Name}}Service = new{{.Name}}Service()

func new{{.Name}}Service() *{{.CamelName}}Service {
	return &{{.CamelName}}Service {}
}

type {{.CamelName}}Service struct {
}

func (s *{{.CamelName}}Service) Get(id int64) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.Get(database.DB(), id)
}

func (s *{{.CamelName}}Service) Take(where ...interface{}) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.Take(database.DB(), where...)
}

func (s *{{.CamelName}}Service) Find(cnd *database.SqlCnd) []model.{{.Name}} {
	return repositories.{{.Name}}Repository.Find(database.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindOne(cnd *database.SqlCnd) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.FindOne(database.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindPageByParams(params *rest.QueryParams) (list []model.{{.Name}}, paging *pagination.Paging) {
	return repositories.{{.Name}}Repository.FindPageByParams(database.DB(), params)
}

func (s *{{.CamelName}}Service) FindPageByCnd(cnd *database.SqlCnd) (list []model.{{.Name}}, paging *pagination.Paging) {
	return repositories.{{.Name}}Repository.FindPageByCnd(database.DB(), cnd)
}

func (s *{{.CamelName}}Service) Count(cnd *database.SqlCnd) int64 {
	return repositories.{{.Name}}Repository.Count(database.DB(), cnd)
}

func (s *{{.CamelName}}Service) Create(t *model.{{.Name}}) error {
	return repositories.{{.Name}}Repository.Create(database.DB(), t)
}

func (s *{{.CamelName}}Service) Update(t *model.{{.Name}}) error {
	return repositories.{{.Name}}Repository.Update(database.DB(), t)
}

func (s *{{.CamelName}}Service) Updates(id int64, columns map[string]interface{}) error {
	return repositories.{{.Name}}Repository.Updates(database.DB(), id, columns)
}

func (s *{{.CamelName}}Service) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.{{.Name}}Repository.UpdateColumn(database.DB(), id, name, value)
}

func (s *{{.CamelName}}Service) Delete(id int64) {
	repositories.{{.Name}}Repository.Delete(database.DB(), id)
}

`))

var controllerTmpl = template.Must(template.New("controller").Parse(`package admin

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/services"
	"github.com/enrichroad/community/database"
	"github.com/enrichroad/community/rest"
	"github.com/kataras/iris/v12"
	"strconv"
)

type {{.Name}}Controller struct {
	Ctx             iris.Context
}

func (c *{{.Name}}Controller) GetBy(id int64) *rest.Resp {
	t := services.{{.Name}}Service.Get(id)
	if t == nil {
		return rest.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return rest.JsonData(t)
}

func (c *{{.Name}}Controller) AnyList() *rest.Resp {
	list, paging := services.{{.Name}}Service.FindPageByParams(rest.NewQueryParams(c.Ctx).PageByReq().Desc("id"))
	return rest.JsonData(&rest.PageResult{Results: list, Page: paging})
}

func (c *{{.Name}}Controller) PostCreate() *rest.Resp {
	t := &model.{{.Name}}{}
	err := rest.ReadForm(c.Ctx, t)
	if err != nil {
		return rest.JsonErrorMsg(err.Error())
	}

	err = services.{{.Name}}Service.Create(t)
	if err != nil {
		return rest.JsonErrorMsg(err.Error())
	}
	return rest.JsonData(t)
}

func (c *{{.Name}}Controller) PostUpdate() *rest.Resp {
	id, err := rest.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return rest.JsonErrorMsg(err.Error())
	}
	t := services.{{.Name}}Service.Get(id)
	if t == nil {
		return rest.JsonErrorMsg("entity not found")
	}

	err = rest.ReadForm(c.Ctx, t)
	if err != nil {
		return rest.JsonErrorMsg(err.Error())
	}

	err = services.{{.Name}}Service.Update(t)
	if err != nil {
		return rest.JsonErrorMsg(err.Error())
	}
	return rest.JsonData(t)
}

`))

var viewIndexTmpl = template.Must(template.New("index.vue").Parse(`
<template>
    <section class="page-container">
        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-form :inline="true" :model="filters">
                <el-form-item>
                    <el-input v-model="filters.name" placeholder="名称"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" v-on:click="list">查询</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleAdd">新增</el-button>
                </el-form-item>
            </el-form>
        </el-col>

        <!--列表-->
        <el-table :data="results" highlight-current-row border v-loading="listLoading"
                  style="width: 100%;" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column prop="id" label="编号"></el-table-column>
            {{range .Fields}}
			<el-table-column prop="{{.CamelName}}" label="{{.CamelName}}"></el-table-column>
			{{end}}
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :page-sizes="[20, 50, 100, 300]"
                           @current-change="handlePageChange"
                           @size-change="handleLimitChange"
                           :current-page="page.page"
                           :page-size="page.limit"
                           :total="page.total"
                           style="float:right;">
            </el-pagination>
        </el-col>

        <!--新增界面-->
        <el-dialog title="新增" :visible.sync="addFormVisible" :close-on-click-modal="false">
            <el-form :model="addForm" label-width="80px" ref="addForm">
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="addForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="addFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
            </div>
        </el-dialog>

        <!--编辑界面-->
        <el-dialog title="编辑" :visible.sync="editFormVisible" :close-on-click-modal="false">
            <el-form :model="editForm" label-width="80px" ref="editForm">
                <el-input v-model="editForm.id" type="hidden"></el-input>
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="editForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="editFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
  export default {
    data() {
      return {
        results: [],
        listLoading: false,
        page: {},
        filters: {},
        selectedRows: [],

        addForm: {},
        addFormVisible: false,
        addLoading: false,

        editForm: {},
        editFormVisible: false,
        editLoading: false,
      }
    },
    mounted() {
      this.list();
    },
    methods: {
		async list() {
		  const params = Object.assign(this.filters, {
			page: this.page.page,
			limit: this.page.limit,
		  })
		  try {
			const data = await this.$axios.post('/api/admin/{{.KebabName}}/list', params)
			this.results = data.results
			this.page = data.page
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  } finally {
			this.listLoading = false
		  }
		},
		async handlePageChange(val) {
		  this.page.page = val
		  await this.list()
		},
		async handleLimitChange(val) {
		  this.page.limit = val
		  await this.list()
		},
		handleAdd() {
		  this.addForm = {
			name: '',
			description: '',
		  }
		  this.addFormVisible = true
		},
		async addSubmit() {
		  try {
			await this.$axios.post('/api/admin/{{.KebabName}}/create', this.addForm)
			this.$message({ message: '提交成功', type: 'success' })
			this.addFormVisible = false
			await this.list()
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
		async handleEdit(index, row) {
		  try {
			const data = await this.$axios.get('/api/admin/{{.KebabName}}/' + row.id)
			this.editForm = Object.assign({}, data)
			this.editFormVisible = true
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
		async editSubmit() {
		  try {
			await this.$axios.post('/api/admin/{{.KebabName}}/update', this.editForm)
			await this.list()
			this.editFormVisible = false
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
	
		handleSelectionChange(val) {
		  this.selectedRows = val
		},
    }
  }
</script>

<style lang="scss" scoped>

</style>

`))

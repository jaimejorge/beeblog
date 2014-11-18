package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["beeblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:CommentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:CommentsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:CommentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:CommentsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:CommentsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:JwtController"] = append(beego.GlobalControllerRouter["beeblog/controllers:JwtController"],
		beego.ControllerComments{
			"IssueToken",
			`/issue-token`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:PostsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PostsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:PostsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PostsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:PostsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PostsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get","options"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:PostsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PostsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:PostsController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PostsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:UsersController"] = append(beego.GlobalControllerRouter["beeblog/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:UsersController"] = append(beego.GlobalControllerRouter["beeblog/controllers:UsersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:UsersController"] = append(beego.GlobalControllerRouter["beeblog/controllers:UsersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:UsersController"] = append(beego.GlobalControllerRouter["beeblog/controllers:UsersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beeblog/controllers:UsersController"] = append(beego.GlobalControllerRouter["beeblog/controllers:UsersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}

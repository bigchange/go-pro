package forms

/**
<form id="user">
    名字：<input name="username" type="text" />
    年龄：<input name="age" type="text" />
    邮箱：<input name="Email" type="text" />
    <input type="submit" value="提交" />
</form>
 */
type UserInfoForm struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}


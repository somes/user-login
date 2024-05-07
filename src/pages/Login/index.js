import './index.scss';
import {Button, Card, Form, Input, message} from 'antd';
import {useDispatch} from "react-redux";
import {fetchLogin} from "@/store/modules/user";
import {useNavigate} from "react-router-dom";

const Login = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const onFinish = async (values) => {
    // 触发异步action fetchLogin
    await dispatch(fetchLogin(values));
    // 跳转到首页
    navigate('/');
    // 提示用户登录成功
    message.success('登录成功');
  }

  return (
    <div className="login">
      <Card className="login-container">
        <h2 className="login-title">Login</h2>
        <Form validateTrigger={['onBlur', 'onChange']} onFinish={onFinish}>
          <Form.Item
            name="username"
            rules={[{required: true, message: '请输入用户名'}]}
          >
            <Input size="large" placeholder="Username"/>
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{required: true, message: '请输入密码'}]}
          >
            <Input size="large" placeholder="Password"/>
          </Form.Item>
          <Form.Item>
            <Button size="large" type="primary" htmlType="submit">Login</Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}

export default Login;

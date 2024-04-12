import './index.scss';
import {Button, Card, Form, Input} from 'antd';

const Login = () => {

  return (
    <div className="login">
      <Card className="login-container">
        <h2 className="login-title">Login</h2>
        <Form validateTrigger={['onBlur', 'onChange']}>
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

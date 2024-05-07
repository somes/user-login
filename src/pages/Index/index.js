import {Layout, Menu, Popconfirm} from 'antd';
import {HomeOutlined, LogoutOutlined,} from '@ant-design/icons';
import './index.scss';
import {Outlet, useLocation, useNavigate} from "react-router-dom";
import {useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";
import {clearUserInfo, fetchUserInfo} from "@/store/modules/user";

const {Header, Sider} = Layout

const items = [
  {
    label: '首页',
    key: '/',
    icon: <HomeOutlined/>,
  },
]

const Index = () => {
  const navigate = useNavigate();
  const onMenuClick = (e) => {
    const path = e.key;
    navigate(path);
  };

  // 获取当前路由路径
  const location = useLocation();

  // 获取用户信息
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchUserInfo());
  }, [dispatch]);
  const username = useSelector(state => state.user.userInfo.username);

  // 退出登录
  const onConfirm = () => {
    dispatch(clearUserInfo());
    navigate('/login');
  };

  return (
    <Layout>
      <Header className="header">
        <div className="logo"/>
        <div className="user-info">
          <span className="user-name">{username}</span>
          <span className="user-logout">
            <Popconfirm title="是否确认退出？" okText="退出" cancelText="取消" onConfirm={onConfirm}>
              <LogoutOutlined/> 退出
            </Popconfirm>
          </span>
        </div>
      </Header>
      <Layout>
        <Sider width={200} className="site-layout-background">
          <Menu
            mode="inline"
            theme="dark"
            style={{height: '100%', borderRight: 0}}
            items={items}
            // defaultSelectedKeys={['/']}
            selectedKeys={[location.pathname]}
            onClick={onMenuClick}
          >
          </Menu>
        </Sider>
        <Layout className="layout-content" style={{padding: 20}}>
          <Outlet/>
        </Layout>
      </Layout>
    </Layout>
  )
}

export default Index;

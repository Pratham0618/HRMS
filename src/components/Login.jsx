import React, { useState } from 'react';
import { UserOutlined, EyeInvisibleOutlined, EyeTwoTone } from '@ant-design/icons';
import { Input, Space, Typography, Button } from 'antd';
import { Navigate } from 'react-router-dom';

const { Title } = Typography;

const LogIn = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);

  const imageUrl = 'https://wallpaper.forfun.com/fetch/8f/8f0b1487338dc0820748ada8adba3df7.jpeg?h=1200&r=0.5'
  const handleLogin = () => {
    if (username === 'admin' && password === 'password') {
      setLoggedIn(true);
    } else {
      alert('Wrong username or password');
    }
  };

  if (loggedIn) {
    return <Navigate to="/dashboard" />;
  }

  return (
    <div style={{ display: 'flex', alignItems: 'stretch', height: '100vh', backgroundColor: 'lightblue' }}>
      <div style={{ flex: 1, overflow: 'hidden' }}>
        <img src={imageUrl} alt="" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
      </div>

      <div style={{ flex: 1, display: 'flex', alignItems: 'center', justifyContent: 'center', padding: '20px' }}>
        <div style={{ 
          display: 'flex', 
          flexDirection: 'column', 
          alignItems: 'center', 
          maxWidth: '400px',
          backgroundColor: 'white', 
          padding: '20px', 
          borderRadius: '8px',
          boxShadow: '0 4px 8px rgba(0,0,0,0.1)'
        }}>
          <Title level={3} style={{ color: 'darkblue', marginBottom: '20px' }}>
            Enter your admin credentials
          </Title>
          <Space direction="vertical" size="large" style={{ width: '100%' }}>
            <Input
              placeholder="Username"
              suffix={<UserOutlined style={{ color: 'darkblue' }} />}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <Input.Password
              placeholder="Password"
              iconRender={(visible) => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Button type="primary" onClick={handleLogin} style={{ width: '100%' }}>Login</Button>
          </Space>
        </div>
      </div>
    </div>
  );
};

export default LogIn;

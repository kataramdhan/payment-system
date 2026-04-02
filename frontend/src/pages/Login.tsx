import { useState } from "react";
import API from "../Api";

type Props = {
  onLogin: () => void;
};


export default function Login({ onLogin }: Props) {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleLogin = async () => {
    const res = await API.post("/login", { email, password });
    localStorage.setItem("token", res.data.token);
    onLogin();
  };

  return (
    <div className="container">
      <div className="card">
        <h2 style={{ marginBottom: 10 }}>🔐 Login</h2>
        <p style={{ color: "#666" }}>Welcome back, please login</p>

        <input
          placeholder="Email"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />

        <button onClick={handleLogin}>Login</button>
      </div>
    </div>
  );
}
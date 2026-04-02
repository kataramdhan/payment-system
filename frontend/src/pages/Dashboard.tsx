import { useEffect, useState } from "react";
import API from "../Api";

type Transaction = {
  id: number;
  amount: number;
  status: string;
};



export default function Dashboard() {
  const [data, setData] = useState<Transaction[]>([]);

  const fetchTransactions = async () => {
    setLoading(true);
    try {
      const res = await API.get<Transaction[]>("/transactions");
      setData(res.data);
    } finally {
      setLoading(false);
    }
  };

  const getStatusClass = (status: string) => {
    if (status === "success") return "badge success";
    if (status === "failed") return "badge failed";
    return "badge pending";
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    window.location.reload();
  };

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetchTransactions();
    const interval = setInterval(fetchTransactions, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="container">
      <div className="card">
        
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <h2>📊 Transactions</h2>
        <button
            style={{ width: "auto", padding: "6px 12px", background: "#ef4444" }}
            onClick={handleLogout}
          >
          Logout
      </button>
</div>
        <p style={{ color: "#666" }}>
          Real-time transaction monitoring
        </p>
        {loading && <p>Loading...</p>}
        <table className="table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Amount</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {data.map((tx) => (
              <tr key={tx.id}>
                <td>{tx.id}</td>
                <td>Rp {tx.amount}</td>
                <td>
                  <span className={getStatusClass(tx.status)}>
                    {tx.status}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
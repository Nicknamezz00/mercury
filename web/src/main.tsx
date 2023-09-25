import {Suspense, useState} from 'react';
import './index.css';
import {Outlet} from "react-router-dom";
import Loading from "@/pages/Loading";

const App = () => {
  const [loading] = useState(true);

  return loading ? (<Loading />) : (
    <Suspense fallback={<Loading />}>
      <Outlet />
    </Suspense>
  );
};

export default App;
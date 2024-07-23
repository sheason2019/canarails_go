import { useNavigate } from '@modern-js/runtime/router';
import { useEffect } from 'react';

export default function Index() {
  const navigate = useNavigate();
  useEffect(() => {
    navigate('/apps');
  }, []);

  return (
    <main>
      <p>Homepage</p>
    </main>
  );
}

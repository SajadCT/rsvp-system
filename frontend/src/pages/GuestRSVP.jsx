import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Calendar, MapPin, Check, X } from 'lucide-react';

const API_URL = 'http://localhost:8080/api';

const GuestRSVP = () => {
  const { id } = useParams();
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(null);

  useEffect(() => {
    axios.get(`${API_URL}/guests/details/${id}`)
      .then(res => { setData(res.data); setStatus(res.data.status); })
      .catch(e => console.error(e))
      .finally(() => setLoading(false));
  }, [id]);

  const handleUpdate = async (newStatus) => {
    try {
      await axios.patch(`${API_URL}/guests/${id}/rsvp`, { status: newStatus });
      setStatus(newStatus);
    } catch (e) { alert("Error updating RSVP"); }
  };

  if (loading) return <div className="h-screen flex items-center justify-center text-gray-500">Loading Invitation...</div>;
  if (!data) return <div className="h-screen flex items-center justify-center text-red-500">Invitation not found.</div>;

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center p-6">
      <div className="bg-white max-w-lg w-full rounded-2xl shadow-xl overflow-hidden">
        <div className="bg-indigo-600 p-8 text-center text-white">
          <h2 className="text-xs font-bold uppercase tracking-widest opacity-80 mb-2">You are invited to</h2>
          <h1 className="text-3xl font-bold mb-4">{data.event_title}</h1>
          <div className="flex justify-center gap-6 text-sm font-medium opacity-90">
             <span className="flex items-center gap-1"><Calendar size={16}/> {data.event_date}</span>
             <span className="flex items-center gap-1"><MapPin size={16}/> {data.event_location}</span>
          </div>
        </div>

        <div className="p-10 text-center">
          <p className="text-xl text-gray-800 mb-2">Hello, <strong>{data.name}</strong>!</p>
          <p className="text-gray-500 mb-8">Please let us know if you can make it.</p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <button onClick={() => handleUpdate('Accepted')}
              className={`flex-1 py-3 px-6 rounded-xl font-bold flex items-center justify-center gap-2 transition ${status === 'Accepted' ? 'bg-green-600 text-white ring-4 ring-green-200' : 'bg-gray-100 text-gray-600 hover:bg-green-50 hover:text-green-600'}`}>
              <Check size={20} /> Going
            </button>
            <button onClick={() => handleUpdate('Declined')}
              className={`flex-1 py-3 px-6 rounded-xl font-bold flex items-center justify-center gap-2 transition ${status === 'Declined' ? 'bg-red-600 text-white ring-4 ring-red-200' : 'bg-gray-100 text-gray-600 hover:bg-red-50 hover:text-red-600'}`}>
              <X size={20} /> Not Going
            </button>
          </div>

          <div className="mt-8 pt-6 border-t border-gray-100 text-sm text-gray-400">
             Current Status: <span className="font-medium text-gray-600">{status}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GuestRSVP;
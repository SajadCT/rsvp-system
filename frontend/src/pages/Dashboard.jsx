import { useState, useEffect } from 'react';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';
import { useNavigate } from 'react-router-dom';
import { 
  Plus, Trash2, MapPin, Calendar, LogOut, Download, 
  Users, CheckCircle, XCircle, Clock, Copy, Check 
} from 'lucide-react';

const API_URL = 'http://localhost:8080/api';


const StatCard = ({ icon: Icon, label, count, color }) => (
  <div className={`p-4 rounded-xl border ${color} bg-white shadow-sm flex items-center gap-4`}>
    <div className={`p-3 rounded-full ${color.replace('border', 'bg').replace('200', '100')}`}>
      <Icon size={24} className={color.replace('border', 'text').replace('200', '600')} />
    </div>
    <div>
      <div className="text-2xl font-bold text-gray-800">{count}</div>
      <div className="text-xs uppercase font-semibold text-gray-500 tracking-wider">{label}</div>
    </div>
  </div>
);

const Dashboard = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  
  const [events, setEvents] = useState([]);
  const [selectedEvent, setSelectedEvent] = useState(null);
  const [guestList, setGuestList] = useState([]);
  const [stats, setStats] = useState({ total: 0, accepted: 0, declined: 0, pending: 0 });

  const [newEvent, setNewEvent] = useState({ title: '', date: '', location: '' });
  const [newGuest, setNewGuest] = useState({ name: '', email: '' });
  const [copiedId, setCopiedId] = useState(null);

  useEffect(() => { fetchEvents(); }, []);
  
  useEffect(() => { if (selectedEvent) fetchDetails(selectedEvent.id); }, [selectedEvent]);

  const fetchEvents = async () => {
    try {
      const res = await axios.get(`${API_URL}/events`);
      setEvents(res.data.data || []);
    } catch (e) { console.error(e); }
  };

  const fetchDetails = async (id) => {
    try {
      const [gRes, sRes] = await Promise.all([
        axios.get(`${API_URL}/guests/${id}`),
        axios.get(`${API_URL}/events/${id}/stats`)
      ]);
      setGuestList(gRes.data.data || []);
      setStats(sRes.data);
    } catch (e) { console.error(e); }
  };

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!newEvent.title) return;
    try {
      await axios.post(`${API_URL}/events`, newEvent);
      setNewEvent({ title: '', date: '', location: '' });
      fetchEvents();
    } catch (e) { alert("Error creating event"); }
  };

  const handleDelete = async (e, id) => {
    e.stopPropagation();
    if (confirm("Delete this event?")) {
      try {
        await axios.delete(`${API_URL}/events/${id}`);
        
        if (selectedEvent?.id === id) setSelectedEvent(null);
        fetchEvents();
      } catch (e) { alert("Error deleting event"); }
    }
  };

  const handleInvite = async (e) => {
    e.preventDefault();
    try {
      
      await axios.post(`${API_URL}/invite`, { ...newGuest, event_id: selectedEvent.id });
      setNewGuest({ name: '', email: '' });
      fetchDetails(selectedEvent.id);
    } catch (e) { alert("Error inviting guest"); }
  };

  const handleExport = async () => {
    try {
      
      const res = await axios.get(`${API_URL}/events/${selectedEvent.id}/export`, { responseType: 'blob' });
      const url = window.URL.createObjectURL(new Blob([res.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', `guests_${selectedEvent.title}.csv`);
      document.body.appendChild(link);
      link.click();
      link.remove();
    } catch (e) { alert("Export failed"); }
  };

  const handleCopyLink = (id) => {
    const link = `${window.location.origin}/rsvp/${id}`;
    navigator.clipboard.writeText(link);
    setCopiedId(id);
    setTimeout(() => setCopiedId(null), 2000);
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col">
      <nav className="bg-white border-b px-8 py-4 flex justify-between items-center shadow-sm sticky top-0 z-10">
        <div className="flex items-center gap-2 text-indigo-600">
          <Calendar size={28} />
          <h1 className="text-xl font-bold tracking-tight">EventFlow</h1>
        </div>
        <div className="flex items-center gap-4">
          <span className="text-gray-600 text-sm font-medium">Hello, {user?.name || 'Host'}</span>
          <button onClick={() => { logout(); navigate('/login'); }} className="flex items-center gap-2 text-sm text-gray-500 hover:text-red-600 transition">
            <LogOut size={16} /> Logout
          </button>
        </div>
      </nav>

      <div className="flex-1 container mx-auto p-6 grid grid-cols-1 lg:grid-cols-12 gap-8">
        
        {}
        <div className="lg:col-span-4 space-y-6">
          <div className="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
            <h2 className="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2">
              <Plus className="text-indigo-500" /> Create Event
            </h2>
            <form onSubmit={handleCreate} className="space-y-3">
              <input className="w-full p-2.5 bg-gray-50 border rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm" placeholder="Event Title" value={newEvent.title} onChange={e => setNewEvent({...newEvent, title: e.target.value})} />
              <div className="flex gap-2">
                <input type="date" className="w-1/2 p-2.5 bg-gray-50 border rounded-lg outline-none text-sm text-gray-600" value={newEvent.date} onChange={e => setNewEvent({...newEvent, date: e.target.value})} />
                <input className="w-1/2 p-2.5 bg-gray-50 border rounded-lg outline-none text-sm" placeholder="Location" value={newEvent.location} onChange={e => setNewEvent({...newEvent, location: e.target.value})} />
              </div>
              <button className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2.5 rounded-lg font-medium transition text-sm">Create Now</button>
            </form>
          </div>

          <div className="space-y-3">
            {events.map(ev => (
            
              <div key={ev.id} onClick={() => setSelectedEvent(ev)}
                className={`group p-5 rounded-xl border cursor-pointer transition-all duration-200 relative ${selectedEvent?.id === ev.id ? 'bg-indigo-50 border-indigo-200 ring-1 ring-indigo-200' : 'bg-white border-gray-100 hover:border-indigo-200 hover:shadow-md'}`}>
                <div className="flex justify-between items-start">
                  <div>
                    <h3 className={`font-bold text-lg ${selectedEvent?.id === ev.id ? 'text-indigo-700' : 'text-gray-800'}`}>{ev.title}</h3>
                    <div className="flex items-center gap-2 text-gray-500 text-xs mt-1">
                      <Calendar size={12} /> {ev.date}
                      <span className="w-1 h-1 bg-gray-300 rounded-full"></span>
                      <MapPin size={12} /> {ev.location}
                    </div>
                  </div>
                  {}
                  <button onClick={(e) => handleDelete(e, ev.id)} className="opacity-0 group-hover:opacity-100 p-2 text-gray-400 hover:text-red-500 transition">
                    <Trash2 size={16} />
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>

        {}
        <div className="lg:col-span-8">
          {selectedEvent ? (
            <div className="space-y-6 fade-in">
              <div className="flex justify-between items-end">
                <div>
                  <h2 className="text-3xl font-bold text-gray-900">{selectedEvent.title}</h2>
                  <p className="text-gray-500 mt-1 flex items-center gap-2"><MapPin size={16}/> {selectedEvent.location}</p>
                </div>
                <button onClick={handleExport} className="flex items-center gap-2 bg-white border border-gray-300 text-gray-700 px-4 py-2 rounded-lg text-sm font-medium hover:bg-gray-50 hover:text-indigo-600 transition shadow-sm">
                  <Download size={16} /> Export CSV
                </button>
              </div>

              <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                <StatCard icon={Users} label="Total" count={stats.total} color="border-gray-200" />
                <StatCard icon={CheckCircle} label="Accepted" count={stats.accepted} color="border-green-200" />
                <StatCard icon={XCircle} label="Declined" count={stats.declined} color="border-red-200" />
                <StatCard icon={Clock} label="Pending" count={stats.pending} color="border-yellow-200" />
              </div>

              <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
                <div className="p-6 border-b border-gray-100 bg-gray-50/50 flex flex-col md:flex-row gap-4 justify-between items-center">
                  <h3 className="font-bold text-gray-700">Guest List</h3>
                  <form onSubmit={handleInvite} className="flex gap-2 w-full md:w-auto">
                    <input className="px-3 py-2 border rounded-lg text-sm outline-none focus:ring-1 focus:ring-indigo-500 w-full" placeholder="Name" value={newGuest.name} onChange={e => setNewGuest({...newGuest, name: e.target.value})} />
                    <input className="px-3 py-2 border rounded-lg text-sm outline-none focus:ring-1 focus:ring-indigo-500 w-full" placeholder="Email" value={newGuest.email} onChange={e => setNewGuest({...newGuest, email: e.target.value})} />
                    <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-indigo-700">Invite</button>
                  </form>
                </div>

                <div className="divide-y divide-gray-100 max-h-125 overflow-y-auto">
                  {guestList.length === 0 ? (
                    <div className="p-10 text-center text-gray-400">No guests invited yet.</div>
                  ) : (
                    guestList.map(guest => (
                      <div key={guest.id} className="p-4 flex justify-between items-center hover:bg-gray-50 transition">
                        <div>
                          <p className="font-semibold text-gray-800">{guest.name}</p>
                          <p className="text-xs text-gray-500">{guest.email}</p>
                        </div>
                        
                        <div className="flex items-center gap-3">
                           <button 
                              onClick={() => handleCopyLink(guest.id)}
                              className={`flex items-center gap-1 text-xs px-2 py-1 rounded border transition
                                ${copiedId === guest.id 
                                  ? 'bg-green-50 border-green-200 text-green-700' 
                                  : 'border-gray-200 text-gray-500 hover:bg-indigo-50 hover:text-indigo-600'
                                }`}
                              title="Copy RSVP Link"
                            >
                              {copiedId === guest.id ? (
                                <><Check size={12} /> Copied</>
                              ) : (
                                <><Copy size={12} /> Copy Link</>
                              )}
                            </button>

                           <span className={`px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide
                            ${guest.status === 'Accepted' ? 'bg-green-100 text-green-700' : 
                              guest.status === 'Declined' ? 'bg-red-100 text-red-700' : 'bg-yellow-100 text-yellow-700'}`}>
                            {guest.status}
                          </span>
                        </div>
                      </div>
                    ))
                  )}
                </div>
              </div>
            </div>
          ) : (
            <div className="h-full flex flex-col items-center justify-center text-gray-400 p-10 border-2 border-dashed border-gray-200 rounded-2xl bg-gray-50">
              <Calendar size={48} className="mb-4 text-gray-300" />
              <p className="text-lg font-medium">Select an event to manage</p>
              <p className="text-sm">Or create a new one on the left</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
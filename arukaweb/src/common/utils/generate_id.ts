// generate id using bcrypt
export function generateId() {
  return crypto.randomUUID();
  // return Date.now().toString() + Math.random().toString();
}

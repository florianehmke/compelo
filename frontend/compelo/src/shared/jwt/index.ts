export function tokenForProjectIdExists(id: string): boolean {
  return projectIdFromToken() === id;
}

export function projectIdFromToken(): string {
  const token = loadToken();
  if (token) {
    const decoded = JSON.parse(atob(token.split('.')[1]));
    return decoded.projectId;
  }
  return "";
}

export function loadToken(): string {
  return localStorage.getItem('compelo-token');
}

export function removeToken() {
  localStorage.removeItem('compelo-token');
}

export function storeToken(token: string) {
  localStorage.setItem('compelo-token', token);
}

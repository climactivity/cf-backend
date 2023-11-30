import { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'de.climatefriday.app',
  appName: 'climatefriday',
  webDir: 'pb_public',
  server: {
    androidScheme: 'https'
  }
};

export default config;

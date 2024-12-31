import React from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Ionicons, FontAwesome5, MaterialCommunityIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import ButtonIcon from '../ButtonIcon';

interface ProviderSelectorProps {
    onProviderSelect: (tech: string) => void;
}

const providers = [
    { name: 'gmail', icon: <MaterialCommunityIcons name="gmail" size={30} color={Colors.light.google} /> },
    { name: 'github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
    { name: 'spotify', icon: <FontAwesome5 name="spotify" size={30} color={Colors.light.spotify} /> },
    { name: 'timer', icon: <MaterialCommunityIcons name="timer" size={30} color={Colors.light.timer} /> },
    { name: 'twitch', icon: <Ionicons name="logo-twitch" size={30} color={Colors.light.twitch} /> },
    { name: 'discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
    { name: 'bitbucket', icon: <FontAwesome5 name="bitbucket" size={30} color={Colors.light.bitbucket} /> },
];

export default function ProviderSelector({ onProviderSelect }: ProviderSelectorProps) {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Select a Provider</Text>
            <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={styles.providersList}>
                {providers.map((tech, index) => (
                    <ButtonIcon
                        key={index}
                        title={tech.name}
                        icon={tech.icon}
                        onPress={() => onProviderSelect(tech.name)}
                        textColor={Colors.light.tint}
                        borderCol={Colors.light.tint}
                        style={{ marginVertical: 5 }}
                    />
                ))}
            </ScrollView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        padding: 20,
        backgroundColor: '#fff',
        borderRadius: 10,
        marginHorizontal: 30,
    },
    title: {
        fontSize: 18,
        color: Colors.light.tint,
        fontWeight: 'bold',
        marginBottom: 15,
        textAlign: 'center',
    },
    providersList: {
        flexDirection: 'column',
    },
});

import json
import os
import hashlib
import re

def sha256_file(filepath):
    sha256 = hashlib.sha256()
    with open(filepath, 'rb') as f:
        for chunk in iter(lambda: f.read(4096), b""):
            sha256.update(chunk)
    return sha256.hexdigest()

def patch_langpacks(base_dir):
    json_path = os.path.join(base_dir, 'official-language-packs.json')
    with open(json_path, 'r', encoding='utf-8') as f:
        data = json.load(f)

    for pack in data.get('packs', []):
        for lang in pack.get('languages', []):
            rel_file = lang.get('file')
            if not rel_file:
                continue
            
            abs_file = os.path.join(base_dir, rel_file)
            if not os.path.exists(abs_file):
                print(f"File not found: {abs_file}")
                continue
                
            # Read and replace content
            with open(abs_file, 'r', encoding='utf-8') as f:
                content = f.read()
                
            # Replace Telegram with Whatsgram (case sensitive)
            new_content = content.replace("Telegram", "Whatsgram")
            new_content = new_content.replace("telegram", "whatsgram")
            new_content = new_content.replace("TELEGRAM", "WHATSGRAM")
            
            # If no changes, skip to avoid unnecessary bumps
            if new_content == content:
                continue
                
            # Increment version
            old_version = lang.get('version', 0)
            new_version = old_version + 1
            lang['version'] = new_version
            
            # Form new filename
            # e.g., android/android_en_v59885528.strings -> android/android_en_v59885529.strings
            new_rel_file = re.sub(r'_v\d+\.strings$', f'_v{new_version}.strings', rel_file)
            
            # If regex didn't match, just append
            if new_rel_file == rel_file:
                new_rel_file = rel_file.replace('.strings', f'_v{new_version}.strings')
                
            new_abs_file = os.path.join(base_dir, new_rel_file)
            
            # Write new file
            with open(new_abs_file, 'w', encoding='utf-8') as f:
                f.write(new_content)
                
            # Remove old file if name changed
            if new_abs_file != abs_file:
                os.remove(abs_file)
                
            # Update json entry
            lang['file'] = new_rel_file
            lang['sha256'] = sha256_file(new_abs_file)
            
            print(f"Updated {rel_file} -> {new_rel_file}")

    # Write back json
    with open(json_path, 'w', encoding='utf-8') as f:
        json.dump(data, f, indent=2, ensure_ascii=False)
        
    print("Done patching langpacks!")

if __name__ == '__main__':
    base_dir = r"c:\Users\myahm\OneDrive\Documents\Whatsgram\gramsrv\data\langpack"
    patch_langpacks(base_dir)

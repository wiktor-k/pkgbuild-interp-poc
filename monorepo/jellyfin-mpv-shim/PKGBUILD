# Maintainer: Giovanni Harting <anonfunc@archlinux.org>
# Contributor: Marius Lindvall <(firstname) {cat} varden {dog} info>

pkgname=jellyfin-mpv-shim
pkgver=3.1.0
pkgrel=1
pkgdesc='Cast media from Jellyfin Mobile and Web apps to MPV'
arch=(any)
url='https://github.com/jellyfin/jellyfin-mpv-shim'
license=(MIT)
depends=(mpv python hicolor-icon-theme python-mpv python-requests python-mpv-jsonipc python-jellyfin-apiclient python-pillow)
makedepends=(python-build python-installer python-wheel gettext python-setuptools)
optdepends=(
  'python-pystray: systray support'
  'mpv-shim-default-shaders: default shader pack'
  'python-pypresence: Discord Rich Presence integration'
)
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz"
  "shaderpack.patch")
b2sums=('157c6a8941d1cdcf52a1ac92bb9c85aa1aebb99cc0db76e87d2ac28792a2fb6112ed77dc4f1b75f270fc3cb8760a4c0df0c7dfe7ad9a46da2ac6f57d16074007'
        '268aca6cb4fb7ad1f3aad5e3a13e34f058f1b17539285b6bdae5b001190bbb123ebcca0a85f83a752dd4fd4d0c80846f80cc7213db5a70f6888c2b13faffa3f7')

prepare() {
  cd jellyfin-mpv-shim-$pkgver

  # remove default-shader-pack from packages
  patch -p1 <../shaderpack.patch
}

build() {
  cd jellyfin-mpv-shim-$pkgver

  find -iname '*.po' | while read -r _file; do
    msgfmt "$_file" -o "${_file%.*}.mo"
  done

  python -m build --wheel --no-isolation
}

package() {
  cd jellyfin-mpv-shim-$pkgver

  install -Dm644 "LICENSE.md" "$pkgdir"/usr/share/licenses/$pkgname/LICENSE

  python -m installer --destdir="$pkgdir" dist/*.whl

  for i in 16 32 48 64 128 256; do
    install -Dvm644 jellyfin_mpv_shim/integration/jellyfin-$i.png "$pkgdir"/usr/share/icons/hicolor/${i}x${i}/apps/com.github.iwalton3.jellyfin-mpv-shim.png
  done

  install -Dm644 jellyfin_mpv_shim/integration/com.github.iwalton3.jellyfin-mpv-shim.desktop -t "$pkgdir"/usr/share/applications/
  install -Dm644 jellyfin_mpv_shim/integration/com.github.iwalton3.jellyfin-mpv-shim.appdata.xml -t "$pkgdir"/usr/share/metainfo/

  cd "$pkgdir"
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  ln -s /usr/share/mpv-shim-default-shaders ${site_packages:1}/jellyfin_mpv_shim/default_shader_pack
}

# vim:set ts=2 sw=2 et:

# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Daniel J Griffiths <ghost1227@archlinux.us>
# Contributor: Ronald van Haren <ronald.archlinux.org>
# Contributor: tobias <tobias@archlinux.org>
# Contributor: Ben <ben@benmazer.net>

pkgname=scribus
pkgver=1.6.6
pkgrel=8
pkgdesc="Desktop publishing software"
arch=(x86_64)
url="https://www.scribus.net/"
license=(GPL-2.0-or-later)
depends=(cairo libcairo.so
         fontconfig libfontconfig.so
         freetype2 libfreetype.so
         glibc # libc.so libm.so
         harfbuzz libharfbuzz.so libharfbuzz-subset.so
         harfbuzz-icu libharfbuzz-icu.so
         hunspell # libhunspell-1.7.so
         icu libicuuc.so
         lcms2 liblcms2.so
         libcdr # libcdr-0.1.so
         libcups # libcups.so
         libfreehand # libfreehand-0.1.so
         libgcc libgcc_s.so
         libjpeg-turbo libjpeg.so
         libmspub # libmspub-0.1.so
         libpagemaker # libpagemaker-0.0.so
         libpng libpng16.so
         libqxp # libqxp-0.0.so
         librevenge # librevenge-0.0.so librevenge-generators-0.0.so librevenge-stream-0.0.so
         libstdc++ libstdc++.so
         libtiff libtiff.so
         libvisio # libvisio-0.1.so
         libxml2 libxml2.so
         libzmf # libzmf-0.0.so
         openscenegraph # libOpenThreads.so libosgDB.so libosgGA.so libosg.so libosgUtil.so libosgViewer.so
         openssl
         podofo libpodofo.so
         poppler libpoppler.so
         python # libpython3.14.so
         qt5-base # libQt5Core.so libQt5Gui.so libQt5Network.so libQt5OpenGL.so libQt5PrintSupport.so libQt5Widgets.so libQt5Xml.so
         zlib libz.so)
makedepends=(cmake
             mesa
             qt5-tools)
optdepends=('gdal: enable gdal plugin'
            'hyphen-lang: hyphenation patterns for desired languages'
            'tk: scripts based on tkinter')
options=(!lto)
_archive="$pkgname-$pkgver"
source=("https://downloads.sourceforge.net/${pkgname}/$_archive.tar.xz"{,.asc}
        "fix_build_with_poppler_26.05.patch::https://github.com/scribusproject/scribus/commit/14a287fc1db2a44abfe1743260554447b31b4adf.patch?full_index=1"
         2b9405a00a96a09e0183190ddc9f83d44963d4e0.patch
        "fix_build_with_poppler_26.07.patch::https://github.com/scribusproject/scribus/commit/ba246c3a2dd086b4e84517723beab159736ba9ba.patch?full_index=1")
sha256sums=('fdfe3e7cbe84b760b38d0561ed8736f9d25d4923adde6e15e03760d83be6166d'
            'SKIP'
            'e84b3f12333cac812c942b85f06403669ad3792bdae821a7e95fa9a3258e1cda'
            'dffd5d9aabd1ae3c8c95832ceb4f3113c5d320ce185770fb961584bbe7bf4bc2'
            '4238d53ae3cb4d545a7c0835a6a0cf082c3b51a3adb74f8cb3dd82977595b4b1')
validpgpkeys=(5086B8D68E70FDDF4C40045AEF7B95E7F60166DA  # Peter Linnell <plinnell@scribus.net>
              757F5E9B13DD648887AD50092D47C099E782504E  # The Scribus Team (www.scribus.net) <the_scribus_team@scribus.net>
              6558BE84D27273A438A151198BEA48118AEBEE64) # Craig Bradney <cbradney@zipworld.com.au>

pkgver() {
        # People regularly flag this OOD pointing to development releases, avoid mistakenly packaging them
	if [[ $pkgver =~ ^[0-9]+\.[13579]+\.[0-9]+$ ]]; then
		echo 'Odd-numbered minor version segments indicate development versions.' >&2
		exit 1
	else
		echo "$pkgver"
	fi
}

prepare() {
	cd "$_archive"
	patch -Np1 -i "$srcdir/fix_build_with_poppler_26.05.patch"
	patch -Np1 --binary -i ../2b9405a00a96a09e0183190ddc9f83d44963d4e0.patch # Fix build with poppler 26.06
	patch -Np1 -i "$srcdir/fix_build_with_poppler_26.07.patch"
}

build() {
	local cmake_options=(
		-D CMAKE_INSTALL_PREFIX=/usr
		-D CMAKE_BUILD_TYPE=None
		-D CMAKE_SKIP_RPATH=On
		-D WANT_CPP20=On
		-D WANT_HUNSPELL=On
		-D WITH_PODOFO=On
		-D WANT_GRAPHICSMAGICK=On
	)
	cmake -B build -S "$_archive" -W no-dev "${cmake_options[@]}"
	make -C build
}

package() {
	make -C build DESTDIR="${pkgdir}" install
	cd "$_archive"
	install -Dm0644 -t "$pkgdir/usr/share/applications/" scribus.desktop
	pushd resources/iconsets/artwork
	for i in 16x16 32x32 128x128 256x256 512x512 1024x1024; do
		install -Dm0644 icon_$i.png "$pkgdir/usr/share/icons/hicolor/$i/apps/scribus.png"
	done
	# Use system hyphen
	rm -rf "$pkgdir/usr/share/scribus/dicts/hyph"
	ln -sf /usr/share/hyphen "$pkgdir/usr/share/scribus/dicts/hyph"
}

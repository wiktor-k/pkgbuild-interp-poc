# Maintainer: Daniel Bermond <dbermond@archlinux.org>

pkgbase=vosk-api
pkgname=(
    'vosk-api'
    'python-vosk')
pkgver=0.3.50
pkgrel=7
_openblas_ver=0.3.20
_clapack_commit=f4e1647de96343fb20eea08a9d4f0e0d04c31c51 # from v3.2.1 branch
_openfst_commit=18e94e63870ebcf79ebb42b7035cd3cb626ec090
_kaldi_commit=bc5baf14231660bd50b7d05788865b4ac6c34481 # from vosk branch
_model_small_ver=0.15
_model_spk_ver=0.4
pkgdesc='Offline speech recognition toolkit'
arch=('x86_64')
url='https://alphacephei.com/vosk/'
license=('Apache-2.0')
makedepends=(
    'cmake'
    'git'
    'gradle'
    'java-environment=21'
    'python'
    'python-build'
    'python-cffi'
    'python-installer'
    'python-requests'
    'python-setuptools'
    'python-srt'
    'python-tqdm'
    'python-websockets'
    'python-wheel')
checkdepends=(
    'ffmpeg'
    'python-numpy')
source=("https://github.com/alphacep/vosk-api/archive/v${pkgver}/${pkgbase}-${pkgver}.tar.gz"
        "https://github.com/xianyi/OpenBLAS/archive/v${_openblas_ver}/openblas-${_openblas_ver}.tar.gz"
        "git+https://github.com/alphacep/clapack.git#commit=${_clapack_commit}"
        "git+https://github.com/alphacep/openfst.git#commit=${_openfst_commit}"
        "kaldi-vosk-g${_kaldi_commit:0:7}.tar.gz"::"https://github.com/alphacep/kaldi/archive/${_kaldi_commit}.tar.gz"
        # models are for running tests in the check() function
        "https://alphacephei.com/vosk/models/vosk-model-small-en-us-${_model_small_ver}.zip"
        "https://alphacephei.com/vosk/models/vosk-model-spk-${_model_spk_ver}.zip"
        '010-vosk-api-gradle9.2-fix.patch')
noextract=("vosk-model-small-en-us-${_model_small_ver}.zip")
sha256sums=('cc1067bcc599c9a2f5f38d4257caf2ac636ba244f7c965cee20293a41024f70f'
            '8495c9affc536253648e942908e88e097f2ec7753ede55aca52e5dead3029e3c'
            'b5653996fc4487fc17d5976e598923abe8ee46973ed581b06312a81c1d56b7a4'
            'cab85e0a4057f1c0dd85bc6f92b0e8437a7ca4859cca4a29482d2fd1d5bd7c4a'
            'be5ee5e9bd2902f43beef8ab1a117525e0813010d660dfe4d22eb7ddee14aede'
            '30f26242c4eb449f948e42cb302dd7a686cb29a3423a8367f99ff41780942498'
            'a74d8f51144484813e16af689bb0f916b7a111e2347f467c4933c1166097b5a7'
            '546cde8eaf5e38de498cf522be7b3990dcf5451c05e1632889441b9c60b28733')

prepare() {
    mkdir -p models
    bsdtar -x -f  "vosk-model-small-en-us-${_model_small_ver}.zip" -C models
    
    ln -sf "../../../vosk-model-spk-${_model_spk_ver}" "${pkgbase}-${pkgver}/python/example/model-spk"
    ln -sf "../../OpenBLAS-${_openblas_ver}" "kaldi-${_kaldi_commit}/tools/OpenBLAS"
    ln -sf ../../clapack "kaldi-${_kaldi_commit}/tools/clapack"
    ln -sf ../../openfst "kaldi-${_kaldi_commit}/tools/openfst"
    
    patch -d "${pkgbase}-${pkgver}" -Np1 -i "${srcdir}/010-vosk-api-gradle9.2-fix.patch"
    
    autoreconf -fi openfst
}

build() {
    local _file
    export CFLAGS+=' -ffat-lto-objects'
    export CXXFLAGS+=' -ffat-lto-objects'
    
    # openblas
    CFLAGS+=' -Wno-implicit-function-declaration' \
    make -C "OpenBLAS-${_openblas_ver}" ONLY_CBLAS='1' DYNAMIC_ARCH='1' TARGET='NEHALEM' USE_LOCKING='1' USE_THREAD='0' all
    make -C "OpenBLAS-${_openblas_ver}" PREFIX="${srcdir}/OpenBLAS-${_openblas_ver}/install" install
    
    # clapack
    CFLAGS+=' -Wno-error=format-security -fcommon -Wno-implicit-function-declaration' \
    cmake -B build-clapack -S clapack \
        -DCMAKE_C_STANDARD:STRING='17' \
        -DCMAKE_POLICY_VERSION_MINIMUM:STRING='3.5.0' \
        -Wno-dev
    cmake --build build-clapack
    while read -r -d '' _file
    do
        cp -af "$_file" "${srcdir}/OpenBLAS-${_openblas_ver}/install/lib"
    done < <(find build-clapack -type f -name 'lib*.a' -print0)
    
    # openfst
    cd openfst
    CFLAGS="${CFLAGS/-O2/-O3}" \
    CXXFLAGS="${CXXFLAGS/-O2/-O3}" \
    ./configure \
        --prefix="${srcdir}/kaldi-${_kaldi_commit}/tools/openfst" \
        --enable-static \
        --enable-shared \
        --enable-far \
        --enable-ngram-fsts \
        --enable-lookahead-fsts \
        --with-pic \
        --disable-bin
    make
    make install
    
    # kaldi
    cd "${srcdir}/kaldi-${_kaldi_commit}/src"
    CXXFLAGS="${CXXFLAGS/-O2/-O3}" ./configure --mathlib='OPENBLAS_CLAPACK' --shared --use-cuda='no'
    sed -i 's/[[:space:]]-O1[[:space:]]/ -O3 /g' kaldi.mk
    make online2 rnnlm
    while read -r -d '' _file
    do
        rm "$_file"
    done < <(find "${srcdir}/kaldi-${_kaldi_commit}" -type f -name '*.o' -print0)
    
    # vosk-api
    export CFLAGS="${CFLAGS/ -ffat-lto-objects/}"
    export CXXFLAGS="${CXXFLAGS/ -ffat-lto-objects/}"
    make -C "${srcdir}/${pkgbase}-${pkgver}/src" \
        EXTRA_CFLAGS="${CXXFLAGS/-O2/-O3}" \
        EXTRA_LDFLAGS="$LDFLAGS" \
        KALDI_ROOT="${srcdir}/kaldi-${_kaldi_commit}" \
        OPENFST_ROOT="${srcdir}/kaldi-${_kaldi_commit}/tools/openfst" \
        OPENBLAS_ROOT="${srcdir}/kaldi-${_kaldi_commit}/tools/OpenBLAS/install"
    
    # java bindings
    gradle -p "${srcdir}/${pkgbase}-${pkgver}/java/lib" clean assemble
    
    # python module
    cd "${srcdir}/${pkgbase}-${pkgver}/python"
    python ./vosk_builder.py
    python -m build --wheel --no-isolation
}

check() {
    local _test
    cd "${pkgbase}-${pkgver}/python/example"
    
    # https://github.com/alphacep/vosk-api/issues/1220
    for _test in alternatives empty ffmpeg nlsml simple speaker srt text words #reset
    do
        printf '%s\n' "Running test_${_test}..."
        PYTHONPATH="${PWD}/../build/lib" VOSK_MODEL_PATH="${srcdir}/models" python "./test_${_test}.py" test.wav
    done
}

package_vosk-api() {
    depends=(
        'glibc'
        'libgcc'
        'libstdc++')
    optdepends=(
        'java-runtime: for java bindings')
    
    install -d -m755 "${pkgdir}/usr/lib"
    install -D -m644 "${pkgbase}-${pkgver}/src/vosk_api.h" -t "${pkgdir}/usr/include"
    install -D -m644 "${pkgbase}-${pkgver}/java/lib/build/libs/vosk-${pkgver}.jar" "${pkgdir}/usr/share/java/vosk.jar"
    cp -dr --no-preserve='ownership' "${pkgbase}-${pkgver}/src"/*.so* "${pkgdir}/usr/lib"
}

package_python-vosk() {
    pkgdesc='Python module for vosk-api'
    depends=(
        'python'
        'python-cffi'
        'python-requests'
        'python-srt'
        'python-tqdm'
        'python-websockets'
        "vosk-api=${pkgver}")
    
    local _site_pkgs
    _site_pkgs="$(python -c 'import site; print(site.getsitepackages()[0])')"
    python -m installer --destdir="$pkgdir" "${pkgbase}-${pkgver}/python/dist"/*.whl
    rm "${pkgdir}${_site_pkgs}/vosk/libvosk.so"
    ln -s ../../../libvosk.so "${pkgdir}${_site_pkgs}/vosk/libvosk.so"
}
